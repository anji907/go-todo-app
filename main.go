package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	// 文字列出力
	fmt.Println("Hello world.")

	app := fiber.New()

	// handle GET request.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello world!"})
	})

	log.Fatal(app.Listen(":4000"))
}
