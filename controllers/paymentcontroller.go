package controllers

import (
	"LightningTipBot/models"
	"github.com/lightningnetwork/lnd/lnrpc"
	"time"
	"LightningTipBot/lnd"
	"LightningTipBot/config"
	"log"
)

func HasPendingWithdrawal(user *models.User) (bool) {
	payments := []models.Payment{}
	err := models.DB.Where("user_id = ?",user.Id).Where("pending = ?",1).Where("success = ?",0).Find(&payments)
	if err.Error != nil {
		log.Print(err.Error)
		return true
	}
	if len(payments) == 0 {
		return false
	}
	return true
}

func CreatePayment(user *models.User, payreq *lnrpc.PayReq, payreqhash string) (*models.Payment, error){

	user.Balance = user.Balance - payreq.NumSatoshis - config.WithdrawalFees
	err := user.Update()
	if err != nil {
		return nil,err
	}

	payment := &models.Payment{
		UserId: user.Id,
		PaymentRequest:payreqhash,
		Amount:payreq.NumSatoshis,
		Success:false,
		Pending:true,
		CreationDate:time.Now(),
	}

	err = payment.Create()
	if err != nil {
		return nil,err
	}

	return payment,nil
}

func ProcessPayment(payment *models.Payment, user *models.User) (*models.Payment,error) {
	pay,err := lnd.SendPaymentSync(payment.PaymentRequest)
	if err != nil {
		payment.Success = false
		payment.Pending = false
		payment.PaymentDate = time.Now()
		payment.Update()

		user.Balance = user.Balance + payment.Amount + config.WithdrawalFees
		user.Update()
		return nil,err
	}

	payment.Fees = pay.PaymentRoute.TotalFees
	payment.Pending = false
	payment.Success = true
	payment.PaymentDate = time.Now()
	payment.Update()

	return payment,nil

}
