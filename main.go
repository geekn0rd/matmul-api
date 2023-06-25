package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Matrix struct {
	Rows int     `json:"rows"`
	Cols int     `json:"cols"`
	Data [][]int `json:"data"`
}

func main() {
	fmt.Println("abbas")
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON("Hello")
	})

	if err := app.Listen("0.0.0.0:1379"); err != nil {
		fmt.Printf("we have error on listening %s\n", err)
	}
}
