package telegram

import (
	"gopkg.in/tucnak/telebot.v2"
	"LightningTipBot/models"
	"LightningTipBot/controllers"
	"strconv"
	"LightningTipBot/lnd"
)

func HelpHandler(b *telebot.Bot, m *telebot.Message, message *models.Message) {

	help_text := ""
	help_text = "Welcome to the lightning tipbot. Start by sending /register to register an account and start using the bot."
	help_text += "\n\nCommands:"
	help_text += "\n\n\\register: Register an account. Make sure you have a telegram username. Your funds are associated with your telegram username so withdraw all your funds if you decide to change your telegram username"
	help_text += "\n\n\\deposit <amount>: Get an invoice address to deposit coins via lightning network"
	help_text += "\n\n\\withdraw <pay_req>: Withdraw your coins over lightning network"
	help_text += "\n\n\\get_node_info: If you get a route not found error while depositing or withdrawing funds you will need to open a channel to our node. This command will provide the node key and address"
	help_text += "\n\n\\tip <amount>: Reply to any message with tip <amount> and the sender of the message will be tipped with the specified amount"

	b.Send(m.Sender, help_text)
	controllers.UpdateResponse(message,help_text)
}


func RegisterHandler(b *telebot.Bot, m *telebot.Message,message *models.Message) {
	if m.Sender.Username != "" {
		user,err := controllers.FindUser(m.Sender.Username)
		if err != nil {
			response := "Some error occurred. Please contact the admin @funyug"
			b.Send(m.Sender,response)
			controllers.UpdateResponse(message,response)
			return
		}

		if user.Id != 0 {
			response := "Already registered"
			b.Send(m.Sender,response)
			controllers.UpdateResponse(message,response)
			return
		}

		user = &models.User{
			Username:m.Sender.Username,
		}
		err = user.Register()
		if err != nil {
			response := "Some error occurred. Please contact the admin @funyug"
			b.Send(m.Sender,response)
			controllers.UpdateResponse(message,response)
		} else {
			response := "Successfully registered"
			b.Send(m.Sender, response)
			controllers.UpdateResponse(message,response)
		}
	} else {
		response := "You need to have a telegram username to register for the bot."
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
	}
}

func DepositHandler(b *telebot.Bot, m *telebot.Message,message *models.Message, payload string) {
	user,err := controllers.FindUser(m.Sender.Username)
	if err != nil {
		response := "Some error occurred. Please contact the admin @funyug"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	if user.Id == 0 {
		response := "You need to be registered to use the bot"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	if payload == "" {
		response := "Amount missing"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	amount, err := strconv.ParseInt(payload,10,64)
	if err != nil {
		response := "Please enter the amount in integers"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	invoice,err := controllers.CreateInvoice(user,amount)
	if err != nil {
		response := "Some error occurred. Please contact the admin @funyug"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
	} else {
		response := invoice.PaymentRequest
		b.Send(m.Sender, response)
		controllers.UpdateResponse(message,response)
	}

}

func WithdrawalHandler(b *telebot.Bot, m *telebot.Message,message *models.Message, payload string) {
	user,err := controllers.FindUser(m.Sender.Username)
	if err != nil {
		response := "Some error occurred. Please contact the admin @funyug"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	if user.Id == 0 {
		response := "You need to be registered to use the bot"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}


	if payload == "" {
		response := "Payment request is missing"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	payreq,err := lnd.DecodePayReq(payload)
	if err != nil {
		response := "Invalid payment request"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	if controllers.HasPendingWithdrawal(user) {
		response := "Please wait for your previous withdrawal to be processed first"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}

	if controllers.HasEnoughBalance(user,payreq) {
		response := "Withdrawal request created"
		b.Send(m.Sender,response)
		controllers.UpdateResponse(message,response)
		return
	}


}