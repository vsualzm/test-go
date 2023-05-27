package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	AddressDate time.Time      `json:"addressDate"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type AddressResponse struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	AddressDate time.Time `json:"addressDate"`
}

func (AddressResponse) TableName() string {
	return "addresss"
}
