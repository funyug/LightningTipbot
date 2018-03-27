package controllers

import (
	"LightningTipBot/models"
	"LightningTipBot/lnd"
	"time"
	"LightningTipBot/config"
	"io"
	"github.com/pkg/errors"
	"log"
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
		invoice.SettledDate = time.Now()
		invoice.Settled = false
		invoice.Create()
		return &invoice,nil
	}
}

func InvoiceSettler() {
	stream,err := lnd.SubscribeInvoices()
	if err != nil {
		config.Fatal(err)
	}
	for {
		i,err := stream.Recv()
		if err == io.EOF {
			new_error := errors.New("Invoice settler stream died")
			config.Fatal(new_error)
			return
		}

		invoice,err := FindInvoice(i.PaymentRequest)
		if err!= nil {
			log.Println(err)
		}
		if invoice.Id != 0 {
			invoice.Settled = i.Settled
			invoice.SettledDate = time.Unix(i.SettleDate, 0)
			invoice.Update()

			user,err := FindUserById(invoice.UserId)
			if err != nil {
				return
			}
			user.Balance = user.Balance + i.Value
			user.Update()

			log.Println("Invoice settled")
		}

	}

}

func FindInvoice(pay_req string) (*models.Invoice,error) {
	i := []models.Invoice{}
	err := models.DB.Where("payment_request = ?",pay_req).Find(&i)
	if err.Error != nil {
		log.Printf("Error occured while finding invoice: %v",err.Error)
		return nil,err.Error
	}
	if len(i) > 0 {
		return &i[0],nil
	} else {
		return &models.Invoice{},nil
	}
}