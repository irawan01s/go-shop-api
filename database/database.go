package database

import (
	"fmt"

	"github.com/irawan01s/go-shop-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.ConfigEnv("DATABASE_USER"),
		config.ConfigEnv("DATABASE_PASSWORD"),
		config.ConfigEnv("DATABASE_HOST"),
		config.ConfigEnv("DATABASE_PORT"),
		config.ConfigEnv("DATABASE_NAME"))

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successfully")
}
