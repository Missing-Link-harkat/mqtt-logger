package dto

import "time"

type TopicDTO struct {
	Topic string `json:"topic"`
}

type SensorDataDTO struct {
	Topic     string    `json:"topic"`
	Value     string    `json:"value"`
	TimeStamp time.Time `json:"timestamp"`
}
