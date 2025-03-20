package main

import (
	"log"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/mqtt"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/utils"
)

func main() {

	utils.LoadEnvVars()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	go mqtt.InitMQTT(dbConn)

	r := api.SetupRouter()
	r.Run(":8080")
}
