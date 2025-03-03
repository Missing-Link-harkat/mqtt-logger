package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/mqtt"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST"),
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
		getEnv("DB_PORT"),
		getEnv("DB_SSLMODE"),
	)
	dbConn, err := db.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	go mqtt.InitMQTT("tcp://localhost:1883", "test/topic", dbConn)

	r := api.SetupRouter()
	r.Run(":8080")
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required env variable: %s", key)
	}
	return value
}
