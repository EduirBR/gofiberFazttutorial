package main

import (
	"fmt"
	"hello-fiber/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()
	// middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello World!") })
	// middleware
	app.Use(requestid.New())
	// user groups
	userGroup := app.Group("/users")
	userGroup.Get("", handlers.HandleUser)
	userGroup.Post("", handlers.HandlerCreateUser)

	app.Listen(":3000")
	fmt.Println("hello World")
}
