package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	// cek error
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/app_transaction?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database")
	}
	fmt.Println("Connected to database")

}
