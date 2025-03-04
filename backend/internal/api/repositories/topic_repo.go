package repositories

import "github.com/Missing-Link-harkat/mqtt-logger/internal/db"

func GetAllTopics() ([]db.Topics, error) {
	var topics []db.Topics
	if err := db.DB.Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}
