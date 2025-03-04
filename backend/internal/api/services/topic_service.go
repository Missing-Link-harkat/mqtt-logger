package services

import (
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

func transformTopicsToDTOs(topics []db.Topics) []dto.TopicDTO {
	var topicDTOs []dto.TopicDTO
	for _, topic := range topics {
		topicDTOs = append(topicDTOs, dto.TopicDTO{Topic: topic.Topic})
	}
	return topicDTOs
}
