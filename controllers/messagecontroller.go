package controllers

import (
	"gopkg.in/tucnak/telebot.v2"
	"LightningTipBot/models"
	"log"
	"time"
)

func StoreMessage(message *telebot.Message) (*models.Message,error) {
	log.Printf("Message received : %v",message.Text)
	m := models.Message{}
	m.Username = message.Sender.Username
	user, err :=  FindUser(m.Username)
	if err != nil {
		return nil,err
	}
	m.UserId = user.Id
	m.Message = message.Text
	m.ReceivedAt = time.Now()
	err = m.Create()
	if err != nil {
		return nil,err
	}
	return &m,nil
}

func UpdateResponse(message *models.Message, response string) {
	message.Response = response
	message.RepliedAt = time.Now()
	message.Update()
	log.Printf("Message replied : %v",response)
}
