package repositories

import (
	"time"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
)

func GetAllTopics() ([]db.Topics, error) {
	var topics []db.Topics
	if err := db.DB.Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func GetSensorDataByTopicAndTime(topic string, startTime time.Time, endTime time.Time) ([]db.Message, error) {
	var sensorData []db.Message
	if err := db.DB.Where("topic = ? AND created_at BETWEEN ? AND ?", topic, startTime, endTime).Find(&sensorData).Error; err != nil {
		return nil, err
	}
	return sensorData, nil
}

func GetLastNSensorDataByTopic(topic string, n int) ([]db.Message, error) {
	var sensorData []db.Message
	if err := db.DB.Where("topic = ?", topic).Order("created_at desc").Limit(n).Find(&sensorData).Error; err != nil {
		return nil, err
	}
	return sensorData, nil
}
