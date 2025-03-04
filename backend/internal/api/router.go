package api

import (
	"github.com/Missing-Link-harkat/mqtt-logger/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies([]string{"localhost"})

	/*r.GET("/messages", getMessages)*/

	r.GET("/topics", handlers.GetTopics)
	return r
}
