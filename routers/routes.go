package routers

import (
	"fiberApi/controllers"

	"github.com/gofiber/fiber/v2"
)

/*
import "github.com/gofiber/fiber/v2/middleware/csrf"
import "github.com/gofiber/fiber/v2/middleware/cors"
import "github.com/gofiber/fiber/v2/middleware/session"
import "github.com/gofiber/fiber/v2/middleware/logger"
*/

// Client Routes should not take care of auth, but it is a small project
func ClientRoute(r fiber.Router) {
	// r.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).SendString("Client Routing Succcess")
	// })
	r.Get("/", controllers.GetAllClients)
}
