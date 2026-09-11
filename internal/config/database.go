package config

import (
	"log"

	"github.com/FerSanti1/expense-tracker-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// Migraciones automáticas
	db.AutoMigrate(&models.User{}, &models.Expense{})
}
