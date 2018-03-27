package telegram

import (
	"gopkg.in/tucnak/telebot.v2"
	"LightningTipBot/controllers"
)

func InitCommands(bot *telebot.Bot) {

	bot.Handle("/help",func(tmessage *telebot.Message) {
			message, err := controllers.StoreMessage(tmessage)
			if err != nil {
				response := "An error occurred. Please contact the admin"
				bot.Send(tmessage.Sender,response)
				return
			}
			HelpHandler(bot,tmessage,message)
	})

	bot.Handle("/register",func(tmessage *telebot.Message) {
		message, err := controllers.StoreMessage(tmessage)
		if err != nil {
			response := "An error occurred. Please contact the admin"
			bot.Send(tmessage.Sender,response)
			return
		}
		RegisterHandler(bot,tmessage,message)
	})

	bot.Handle("/deposit",func(tmessage *telebot.Message) {
		message, err := controllers.StoreMessage(tmessage)
		if err != nil {
			response := "An error occurred. Please contact the admin"
			bot.Send(tmessage.Sender,response)
			return
		}
		DepositHandler(bot,tmessage,message,tmessage.Payload)
	})

	bot.Handle("/withdraw",func(tmessage *telebot.Message) {
		message, err := controllers.StoreMessage(tmessage)
		if err != nil {
			response := "An error occurred. Please contact the admin"
			bot.Send(tmessage.Sender,response)
			return
		}
		WithdrawalHandler(bot,tmessage,message,tmessage.Payload)
	})

}
