package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("hello worldsssss")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/",func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"Hello World"})
	})

	//create a TODO
	app.Post("/api/todos",func(c *fiber.Ctx) error {
		todo := &Todo{}

		
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error":"body is required"})
		}
		todo.ID = len(todos) + 1
		todos =append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

	//Update a TODO
	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")
		
		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})

	log.Fatal(app.Listen(":3000"))
}
