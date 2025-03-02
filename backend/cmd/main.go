package main

import (
	"github.com/Missing-Link-harkat/mqtt-logger/internal/api"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/mqtt"
)


func main() {


	go mqtt.InitMQTT("tcp://localhost:1883", "test/topic")


	dsn := "host=localhost user=postgres password=example dbname=example_db port=5432 sslmode=disable"
	go db.ConnectDB(dsn)

	r := api.SetupRouter()
	r.Run(":8080")
}