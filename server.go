package main

import (
	"google.golang.org/grpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"LightningTipBot/lnd"
	"gopkg.in/tucnak/telebot.v2"
	"time"
	"LightningTipBot/config"
	"log"
	"LightningTipBot/telegram"
	"LightningTipBot/models"
	"LightningTipBot/controllers"
)

func main() {
	config.CheckFlags()

	err := models.InitDB()
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	defer models.DB.Close()

	var conn *grpc.ClientConn

	conn = lnd.Connect(conn)
	defer conn.Close()

	lnd.Client = lnrpc.NewLightningClient(conn)

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		config.Fatal(err)
	}

	telegram.InitCommands(bot)
	go controllers.InvoiceSettler()

	log.Println("Server started..")
	bot.Start()

}


