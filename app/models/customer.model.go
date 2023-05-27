package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           int               `json:"id" gorm:"primaryKey"`
	CustomerName string            `json:"customer_name" gorm:"not null"`
	Address      []AddressResponse `json:"customer_address"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdateAt     time.Time         `json:"update_at"`
	DeletedAt    gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

type CustomerResponse struct {
	ID           int               `json:"id" gorm:"primaryKey"`
	CustomerName string            `json:"customer_name" gorm:"not null"`
	Address      []AddressResponse `json:"customer_address"`
}

func (CustomerResponse) TableName() string {
	return "transactions"
}
