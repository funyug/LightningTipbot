package models

import (
	"time"
	"log"
)

type Payment struct {
	Id int
	UserId int
	PaymentRequest string
	Amount int64
	Fees int64
	Pending bool
	Success bool
	CreationDate time.Time
	PaymentDate time.Time
}

func(p *Payment) Create() error {
	err := DB.Create(p)
	if err.Error != nil {
		log.Printf("Error occured while creating payment: %v",err)
	}
	return err.Error
}

func(p *Payment) Update() error {
	err := DB.Save(p)
	if err.Error != nil {
		log.Printf("Error occured while updating payment: %v",err)
	}
	return err.Error
}