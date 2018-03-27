package models

import (
	"time"
	"log"
)

type Invoice struct {
	Id int
	UserId int
	Amount int64
	PaymentRequest string
	Memo string
	Settled bool
	CreationDate time.Time
	SettledDate time.Time
}

func (i *Invoice) Create() error {
	err := DB.Create(i)
	if err.Error != nil {
		log.Printf("Error occured while creating invoice: %v",err)
	}
	return err.Error
}