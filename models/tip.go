package models

import "log"

type Tip struct {
	Id int
	FromUserId int
	ToUserId int
	Amount int64
}

func (t *Tip) Create() error {
	err := DB.Create(t)
	if err.Error != nil {
		log.Printf("Error occured while creating tip: %v",err)
	}
	return err.Error
}