package handlers

import (
	"net/http"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api/services"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {
	topics, err := services.FetchTopics()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	if len(topics) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No topics found"})
		return
	}

	c.JSON(http.StatusOK, topics)
}
