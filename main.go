package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FerSanti1/expense-tracker-api/internal/config"
	"github.com/FerSanti1/expense-tracker-api/internal/middleware"
	"github.com/FerSanti1/expense-tracker-api/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET not set")
	}

	config.InitDB()
	log.Println("Database connected")

	jwtManager := middleware.NewJWTManager(jwtSecret, 10*time.Minute)
	if jwtManager == nil {
		panic("Failed to create JWT manager")
	}

	e := echo.New()

	routes.SetupRoutes(e)
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
