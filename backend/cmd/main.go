package main

import (
	"mqtt--logger/internal/api"
	"mqtt--logger/internal/mqtt"
)


func main() {


	mqtt.InitMQTT("tcp://localhost:1883", "test/topic")

	r := api.SetupRouter()
	r.Run(":8080")
}