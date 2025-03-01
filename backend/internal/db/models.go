package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Topic string
	Payload string
}

type Log struct {
	gorm.Model
	Message string
}