package controllers

import (
	"github.com/FerSanti1/expense-tracker-api/internal/models"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func GetExpenses() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var db *gorm.DB
		var expenses []models.Expense
		db.Find(&expenses)
		return c.JSON(200, expenses)
	}
}
