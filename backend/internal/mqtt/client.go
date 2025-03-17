package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

var mqttClient mqtt.Client

func InitMQTT(dbConn *gorm.DB) {

	mqttBroker := fmt.Sprintf("tcp://%s:%s", utils.GetEnv("MQTT_HOST"), utils.GetEnv("MQTT_PORT"))

	mqttUsername := utils.GetEnv("MQTT_USERNAME")
	mqttPassword := utils.GetEnv("MQTT_PASSWORD")
	mqttCACertPath := utils.GetEnv("MQTT_CA_CERT_PATH")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("go_mqtt_client")

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(3 * time.Second)
	// Auth
	opts.SetUsername(mqttUsername)
	opts.SetPassword(mqttPassword)

	// SSL/TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	if mqttCACertPath != "" {
		caCert, err := os.ReadFile(mqttCACertPath)
		if err != nil {
			log.Fatalf("Failed to read CA certificate: %v", err)
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			log.Fatalf("Failed to append CA certificate to CertPool")
		}
		tlsConfig.RootCAs = certPool
	}

	opts.SetTLSConfig(tlsConfig)

	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("MQTT connection lost: %v", err)
	})

	opts.OnConnect = func(client mqtt.Client) {
		log.Println("Connected to MQTT broker")

		// Retrieve topics
		var topics []db.Topics
		if err := dbConn.Find(&topics).Error; err != nil {
			log.Printf("Error retrieving topics from db: %v", err)
			return
		}

		for _, topic := range topics {
			if token := client.Subscribe(topic.Topic, 1, func(client mqtt.Client, msg mqtt.Message) {
				messageHandler(client, msg, dbConn)
			}); token.Wait() && token.Error() != nil {
				log.Printf("Subscription error: %v", token.Error())
			}
			log.Printf("Subscribed to: %v", topic.Topic)
		}

	}

	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection error: %v", token.Error())
	}
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
