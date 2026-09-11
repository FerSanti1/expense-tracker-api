package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID   uint
	Amount   float64
	Category string
	Date     string
}
