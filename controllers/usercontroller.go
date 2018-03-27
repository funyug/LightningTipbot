package controllers

import (
	"LightningTipBot/models"
	"log"
	"github.com/lightningnetwork/lnd/lnrpc"
)

func FindUser(username string) (*models.User,error) {
	u := []models.User{}
	err := models.DB.Where("username = ?",username).Find(&u)
	if err.Error != nil {
		log.Printf("Error occured while finding user: %v",err.Error)
		return nil,err.Error
	}
	if len(u) > 0 {
		return &u[0],nil
	} else {
		return &models.User{},nil
	}
}


func FindUserById(id int) (*models.User,error) {
	u := []models.User{}
	err := models.DB.Where("id = ?",id).Find(&u)
	if err.Error != nil {
		log.Printf("Error occured while finding user by id: %v",err.Error)
		return nil,err.Error
	}
	if len(u) > 0 {
		return &u[0],nil
	} else {
		return &models.User{},nil
	}
}

func HasEnoughBalance(user *models.User, req *lnrpc.PayReq) bool {
	if user.Balance > req.NumSatoshis + 10 {
		return true
	}
	return false
}
