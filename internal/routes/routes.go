package routes

import (
	"github.com/FerSanti1/expense-tracker-api/internal/controllers"
	"github.com/labstack/echo/v5"
)

func SetupRoutes(e *echo.Echo) {
	expenses := e.Group("/v1/expenses")
	expenses.GET("/", controllers.GetExpenses())
	expenses.POST("/", controllers.AddExpense())
	auth := e.Group("/v1/auth")
	auth.POST("/login", controllers.Login())
	auth.POST("/signup", controllers.SignUp())
}
