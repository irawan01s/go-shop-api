package database

import (
	"fmt"

	"github.com/irawan01s/go-shop-api/models"
)

func InitMigarate() {
	DB.AutoMigrate(&models.User{})

	fmt.Println("Migration success")
}
