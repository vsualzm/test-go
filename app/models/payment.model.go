package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type PaymentResponse struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (PaymentResponse) TableName() string {
	return "payments"
}
