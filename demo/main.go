package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
