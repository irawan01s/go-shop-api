package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/irawan01s/go-shop-api/database"
	"github.com/irawan01s/go-shop-api/routes"
)

func main() {
	database.ConnectDB()
	database.InitMigarate()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8081"))
}
