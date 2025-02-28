package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMessages(c *gin.Context) {
	var message = "hey there"
	c.JSON(http.StatusOK, message)
}