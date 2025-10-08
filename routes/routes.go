package routes

import (
	"github.com/corp-ais/golang-fiber-template/internal/ports"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router, contentPort ports.HandlerPort) {
	// 	healthGroup := app.Group("health")
	// 	healthGroup.Get("/liveness", func(c *fiber.Ctx) error {
	// 		return c.SendStatus(fiber.StatusOK)
	// 	})
	// 	healthGroup.Get("/readiness", func(c *fiber.Ctx) error {
	// 		return c.SendStatus(fiber.StatusOK)
	// 	})

	// ikmContentGroup := app.Group("content", middleware.DecryptMiddleware())
	// ikmContentGroup.Get("/:contentId", contentPort.GetContentHandler)
}
