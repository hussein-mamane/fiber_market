package main

import (
	"fiberApi/routers"

	"github.com/gofiber/fiber/v2"
)

func init() {

}
func main() {
	app := fiber.New(fiber.Config{})
	app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"Message": "Well Done !"})
	})
	// client_handler := app.Group("/clients")
	routers.ClientRoute(app.Group("/clients"))

	app.Listen(":3000")

}
