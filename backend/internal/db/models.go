package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Topic   string
	Payload string
}

type Topics struct {
	gorm.Model
	Topic string
}

type Log struct {
	gorm.Model
	Message string
}
