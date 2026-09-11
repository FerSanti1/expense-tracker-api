package controllers

import (
	"errors"

	"github.com/FerSanti1/expense-tracker-api/internal/models"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var db *gorm.DB
		var User struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.Bind(&User); err != nil {
			return errors.New("Invalid request")
		}
		if err := db.Create(&User).Error; err != nil {
			return errors.New("Failed to create user")
		}
		// we hash the password before saving it to the database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("Failed to hash password")
		}
		User.Password = string(hashedPassword)

		user := models.User{Name: User.Name, Email: User.Email, Password: User.Password}
		if err := db.Create(&user).Error; err != nil {
			return errors.New("Failed to save user")
		}

		return c.JSON(200, user)
	}
}

func Login() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var db *gorm.DB
		var user models.User
		var User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.Bind(&User); err != nil {
			return errors.New("Invalid request")
		}
		if err := db.Where("email = ?", User.Email).First(&user).Error; err != nil {
			return errors.New("Invalid credentials")
		}
		// we compare the hashed password with the plain text password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(User.Password)); err != nil {
			return errors.New("Unmatched password")
		}
		return c.JSON(200, user)
	}
}

func AddExpense() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var db *gorm.DB
		var expense models.Expense
		if err := c.Bind(&expense); err != nil {
			return errors.New("Invalid request")
		}
		if err := db.Create(&expense).Error; err != nil {
			return errors.New("Failed to save expense")
		}
		return c.JSON(200, expense)
	}
}
