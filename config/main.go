package config

import (
"flag"
"errors"
"log"
)

var Token string;

func CheckFlags() {
	var tokenPtr = flag.String("token","","Your telegram bot token")
	flag.Parse()

	Token = *tokenPtr
	if Token == "" {
		err := errors.New("flag token is missing")
		Fatal(err)
	}

}

func Fatal(err error) {
	log.Fatalf( "[lncli] %v\n", err)
}

