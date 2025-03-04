package services

import (
	"errors"
	"time"

	"github.com/Missing-Link-harkat/mqtt-logger/internal/api/dto"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/api/repositories"
	"github.com/Missing-Link-harkat/mqtt-logger/internal/db"
)

func FetchTopics() ([]dto.TopicDTO, error) {
	topics, err := repositories.GetAllTopics()
	if err != nil {
		return nil, err
	}
	return transformTopicsToDTOs(topics), nil
}

func FetchSensorData(topic string, startTime string, endTime string) ([]dto.SensorDataDTO, error) {

	topics, err := repositories.GetAllTopics()
	if err != nil {
		return nil, err
	}
	var topicExists bool
	for _, t := range topics {
		if t.Topic == topic {
			topicExists = true
			break
		}
	}

	if !topicExists {
		return nil, errors.New("topic not found")
	}

	var sensorData []db.Message

	if (startTime != "" && endTime == "") || (startTime == "" && endTime != "") {
		return nil, errors.New("both start_time and end_time must be specified together")
	}

	if startTime != "" && endTime != "" {
		start, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			return nil, errors.New("invalid start time format")
		}
		end, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			return nil, errors.New("invalid end time format")
		}
		sensorData, err = repositories.GetSensorDataByTopicAndTime(topic, start, end)
	} else {
		sensorData, err = repositories.GetLastNSensorDataByTopic(topic, 20)
	}

	if err != nil {
		return nil, err
	}

	return transformSensorDataToDTOs(sensorData), nil
}

func transformTopicsToDTOs(topics []db.Topics) []dto.TopicDTO {
	var topicDTOs []dto.TopicDTO
	for _, topic := range topics {
		topicDTOs = append(topicDTOs, dto.TopicDTO{Topic: topic.Topic})
	}
	return topicDTOs
}

func transformSensorDataToDTOs(sensorData []db.Message) []dto.SensorDataDTO {
	var sensorDataDTOs []dto.SensorDataDTO
	for _, data := range sensorData {
		sensorDataDTOs = append(sensorDataDTOs, dto.SensorDataDTO{
			Topic:     data.Topic,
			Value:     data.Payload,
			TimeStamp: data.CreatedAt,
		})
	}
	return sensorDataDTOs
}
