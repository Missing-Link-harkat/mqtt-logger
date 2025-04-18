package api

import (
	"github.com/Missing-Link-harkat/mqtt-logger/internal/api/handlers"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(utils.GetEnv("GIN_MODE"))
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{utils.GetEnv("ALLOWED_ORIGIN")},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/topics", handlers.GetTopics)
	r.GET("topics/data", handlers.GetSensorData)
	return r
}
