package mqtt

import (
	"fmt"
	"log"
	"time"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

var mqttClient mqtt.Client

func InitMQTT(broker string, topic string, dbConn *gorm.DB) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go_mqtt_client")

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(3 * time.Second)

	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("MQTT connection lost: %v", err)
	})

	opts.OnConnect = func(client mqtt.Client) {
		log.Println("Connected to MQTT broker")
		if token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
			messageHandler(client, msg, dbConn)
		}); token.Wait() && token.Error() != nil {
			log.Printf("Subscription error: %v", token.Error())
		}
	}

	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection error: %v", token.Error())
	}

	mqttClient.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		messageHandler(client, msg, dbConn)
	})
	fmt.Println("Subscribed to: ", topic)
}

func messageHandler(client mqtt.Client, msg mqtt.Message, dbConn *gorm.DB) {
	fmt.Printf("Received message on %s: %s\n", msg.Topic(), string(msg.Payload()))

	topic := msg.Topic()
	payload := string(msg.Payload())

	message := db.Message{
		Topic:   topic,
		Payload: payload,
	}
	if err := dbConn.Create(&message).Error; err != nil {
		log.Printf("Failed to save message to db: %v", err)
	} else {
		log.Printf("Saved message to db")
	}
}
