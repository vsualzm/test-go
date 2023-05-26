package migration

import (
	"fmt"
	"log"
	"test-go/database"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&entity.Transaction{},
		&entity.TransactionDetail{},
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("database berhasil migration")

}
