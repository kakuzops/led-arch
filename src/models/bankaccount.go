package models

import (
	"time"

	"gorm.io/gorm"
)

type BankAccount struct {
	gorm.Model
	Type      string
	Name      string
	Balance   uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
