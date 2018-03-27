package controllers

import (
	"LightningTipBot/models"
	"log"
	"github.com/lightningnetwork/lnd/lnrpc"
)

func FindUser(username string) (*models.User,error) {
	u := []models.User{}
	err := models.DB.Find(&u).Where("username = ?",username)
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

func HasEnoughBalance(user *models.User, req *lnrpc.PayReq) bool {
	if user.Balance > req.NumSatoshis {
		return true
	}
	return false
}
