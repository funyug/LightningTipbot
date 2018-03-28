package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"LightningTipBot/config"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open("mysql", config.DBUser+":"+config.DBPassword+"@/"+config.DBName+"?parseTime=true")
	if err != nil {
		return err
	}
	return nil
}

