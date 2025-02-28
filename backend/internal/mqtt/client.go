package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func InitMQTT(broker string, topic string) {
    opts := mqtt.NewClientOptions()
    opts.AddBroker(broker)
    opts.SetClientID("go_mqtt_client")

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