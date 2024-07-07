package controllers

import (
	"fiberApi/models"

	"github.com/gofiber/fiber/v2"
)

var clients = []*models.Client{
	{
		ClientId:  1,
		FirstName: "Bob",
		LastName:  "Moran",
		Email:     "Moranbob@mail.com",
	},
}

func GetAllClients(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusOK).SendString("Client Routing OK")
	return c.JSON(clients)

}
