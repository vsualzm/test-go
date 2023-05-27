package migrations

import (
	"app-transaction/database"
	"app-transaction/models"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.Customer{},
		&models.Product{},
		&models.Payment{},
		&models.Address{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println("Can't running migration")
	}

	fmt.Println("Migration Success")
}
