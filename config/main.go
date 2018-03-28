package config

import (
	"flag"
	"errors"
	"log"
)

var Token string;
var WithdrawalFees = int64(10);
var DBUser string;
var DBPassword string;
var DBName string;

func CheckFlags() {
	var tokenPtr = flag.String("token","","Your telegram bot token")
	var dbuserPtr = flag.String("dbuser","","Your dbusername")
	var dbpassPtr = flag.String("dbpass","","Your dbpassword")
	var dbnamePtr = flag.String("dbname","","Your dbname")
	flag.Parse()

	Token = *tokenPtr
	DBUser = *dbuserPtr
	DBPassword = *dbpassPtr
	DBName = *dbnamePtr
	if Token == "" {
		err := errors.New("flag token is missing")
		Fatal(err)
	}

	if DBUser == "" {
		err := errors.New("flag dbuser is missing")
		Fatal(err)
	}

	if DBPassword == "" {
		err := errors.New("flag dbpass is missing")
		Fatal(err)
	}

	if DBName == "" {
		err := errors.New("flag dbname is missing")
		Fatal(err)
	}

}

func Fatal(err error) {
	log.Fatalf( "[lncli] %v\n", err)
}

