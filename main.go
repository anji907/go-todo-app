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

	todos := []Todo{}

	// handle GET request.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello world!"})
	})

	// handle POST request.
	app.Post("api/todos", func(c *fiber.Ctx) error {
		// request body.
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(4000).JSON(fiber.Map{"error": "Todo Body is required."})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	log.Fatal(app.Listen(":4000"))
}
