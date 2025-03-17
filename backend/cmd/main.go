package main

import (
	"fmt"
	"log"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/mqtt"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/utils"
)

func main() {

	utils.LoadEnvVars()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_USER"),
		utils.GetEnv("DB_PASSWORD"),
		utils.GetEnv("DB_NAME"),
		utils.GetEnv("DB_PORT"),
		utils.GetEnv("DB_SSLMODE"),
	)
	dbConn, err := db.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	go mqtt.InitMQTT(dbConn)

	r := api.SetupRouter()
	r.Run(":8080")
}
