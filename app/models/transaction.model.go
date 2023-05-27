package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        int               `json:"id" form:"id" gorm:"primaryKey"`
	Product   []Product         `json:"product"`
	Customer  CustomerResponse  `json:"posts"`
	Payment   []PaymentResponse `json:"payment"`
	Address   AddressResponse   `json:"address"`
	CreatedAt time.Time         `json:"created_at"`
	UpdateAt  time.Time         `json:"update_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}
