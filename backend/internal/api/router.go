package api

import "github.com/gin-gonic/gin"


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/messages", getMessages)
	return r
}