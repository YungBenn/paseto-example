package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type login struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

func main() {
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		var input login
		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		token, payload, err := CreateToken(input.Username, input.FullName, 3)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"user":  payload,
			"token": token,
		})
	})

	app.Get("/profile", Authenticate(), func(c *fiber.Ctx) error {
		payload := GetCurrentUser(c)

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"user": payload,
		})
	})

	app.Listen(":8002")
}
