package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irawan01s/go-shop-api/controllers"
)

func UserRoutes(router fiber.Router) {
	user := router.Group("/users")

	user.Get("/", controllers.GetUser)

	user.Post("/", controllers.CreateUser)
}
