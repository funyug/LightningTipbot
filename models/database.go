package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open("mysql", "root:@/lightning?parseTime=true")
	if err != nil {
		return err
	}
	return nil
}

