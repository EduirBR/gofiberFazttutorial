package handlers

import (
	"hello-fiber/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleUser(c *fiber.Ctx) error {
	user := models.User{
		Firstname: "Joe",
		Lastname:  "whillier",
	}
	return c.Status(fiber.StatusAccepted).JSON(user)
}

func HandlerCreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}
