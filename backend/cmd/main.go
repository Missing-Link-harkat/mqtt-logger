package main

import (
	"log"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/mqtt"
)

func main() {

	dsn := "host=localhost user=postgres password=example dbname=example_db port=5432 sslmode=disable"
	dbConn, err := db.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	go mqtt.InitMQTT("tcp://localhost:1883", "test/topic", dbConn)

	r := api.SetupRouter()
	r.Run(":8080")
}
