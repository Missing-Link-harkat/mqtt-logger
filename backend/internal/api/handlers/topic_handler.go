package handlers

import (
	"log"
	"net/http"
	"net/url"

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

func GetSensorData(c *gin.Context) {
	topic := c.DefaultQuery("topic", "")

	/*
		TODO: Proper error handling
	*/
	if topic == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Topic is required"})
		return
	}

	decodedTopic, err := url.QueryUnescape(topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode topic"})
		return
	}
	log.Printf("%v", decodedTopic)

	startTime := c.DefaultQuery("start_time", "")
	endTime := c.DefaultQuery("end_time", "")

	data, err := services.FetchSensorData(decodedTopic, startTime, endTime)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}
