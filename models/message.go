package models

import (
	"log"
	"time"
)

type Message struct {
	Id int
	UserId int
	Username string
	Message string
	Response string
	ReceivedAt time.Time
	RepliedAt time.Time
}

func(m *Message) Create() error {
	err := DB.Create(m)
	if err.Error != nil {
		log.Printf("Error occured while creating message: %v",err)
	}
	return err.Error
}

func(m *Message) Update() error {
	err := DB.Save(m)
	if err.Error != nil {
		log.Printf("Error occured while updating message: %v",err)
	}
	return err.Error
}
