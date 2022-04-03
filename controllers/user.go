package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/irawan01s/go-shop-api/database"
	"github.com/irawan01s/go-shop-api/helper"
	"github.com/irawan01s/go-shop-api/models"
)

type userResponse struct {
	ID       uuid.UUID
	Name     string
	Email    *string
	Phone    string
	Password string
	Address  string
	Birthday string
	Status   bool
	Products []models.Product
}

func GetUser(c *fiber.Ctx) error {
	var users []models.User

	db := database.DB

	db.Preload("Products").Find(&users)
	if len(users) == 0 {
		resError := helper.ErrorResponse("User is empty", "empty", nil)
		return c.Status(fiber.StatusNotFound).JSON(resError)
	}

	list := []userResponse{}
	for _, u := range users {
		birhthDayFormatted := u.Birthday.Format("2006-01-02")
		fmt.Println(birhthDayFormatted)

		list = append(list, userResponse{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Phone:    u.Phone,
			Address:  u.Address,
			Birthday: birhthDayFormatted,
			Status:   u.Status,
			Products: u.Products,
		})
	}

	// format := "2006-01-02"
	// date, _ := time.Parse(format, "2019-07-10")
	// fmt.Println(date)

	resOk := helper.SuccessResponse(true, "Success", list)
	return c.Status(fiber.StatusOK).JSON(resOk)
}

func CreateUser(c *fiber.Ctx) error {
	var err error
	var resError helper.Response

	db := database.DB
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		resError = helper.ErrorResponse("Bad request", err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(resError)
	}

	user.ID = uuid.New()

	err = db.Create(&user).Error
	if err != nil {
		resError = helper.ErrorResponse("Fail created", err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(resError)
	}

	resOk := helper.SuccessResponse(true, "Success", user)

	return c.Status(fiber.StatusOK).JSON(resOk)
}
