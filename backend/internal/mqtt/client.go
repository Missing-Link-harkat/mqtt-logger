package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func InitMQTT(broker string, topic string) {
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
        if token := client.Subscribe(topic, 1, messageHandler); token.Wait() && token.Error() != nil {
            log.Printf("Subscription error: %v", token.Error())
        }
    }

    mqttClient = mqtt.NewClient(opts)
    if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("MQTT connection error: %v", token.Error())
    }

    mqttClient.Subscribe(topic, 1, messageHandler)
    fmt.Println("Subscribed to: ", topic)
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message on %s: %s\n", msg.Topic(), string(msg.Payload()))
    
}