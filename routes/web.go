package routes

import "github.com/gofiber/fiber/v2"

func LoadRoutes(r fiber.Router) {
	// Test to load static, compiled assets
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Golang   s:)")
	})
	r.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("YOYOYOYO :)")
	})
}
