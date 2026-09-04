package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
