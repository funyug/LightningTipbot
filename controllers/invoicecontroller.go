package controllers

import (
	"LightningTipBot/models"
	"LightningTipBot/lnd"
	"time"
)

func CreateInvoice(user *models.User, amount int64) (*models.Invoice,error){
	invoice := models.Invoice{
		UserId:user.Id,
		Amount:amount,
	}

	response,err := lnd.AddInvoice(amount)
	if err != nil {
		return nil,err
	} else {
		invoice.PaymentRequest = response.PaymentRequest
		invoice.CreationDate = time.Now()
		invoice.Settled = false
		invoice.Create()
		return &invoice,nil
	}
}
