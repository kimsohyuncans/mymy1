package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mejik-dev/microgen-v3-go"
)

var API_KEY = os.Getenv("API_KEY")

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Wordfld!")
	})

	client := microgen.NewClient(API_KEY, microgen.DefaultURL())

	products := app.Group("/products")

	products.Get("", func(c *fiber.Ctx) error {
		resp, err := client.Service("products").Find()
		if err != nil {
			if err.Message == "project not found" {
				return c.Status(err.Status).JSON(map[string]interface{}{
					"message": "failed to connect to your project, please check if the api had been set properly.",
				})
			}

			return c.Status(err.Status).JSON(err)
		}

		return c.Status(http.StatusOK).JSON(resp.Data)
	})

	products.Post("", func(c *fiber.Ctx) error {
		body := make(map[string]interface{})

		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(http.StatusBadRequest).SendString("failed parse request body to json")
		}

		resp, err := client.Service("products").Create(body)
		if err != nil {
			if err.Message == "project not found" {
				return c.Status(err.Status).JSON(map[string]interface{}{
					"message": "failed to connect to your project, please check if the api had been set properly.",
				})
			}

			return c.Status(err.Status).JSON(err)
		}

		return c.Status(http.StatusOK).JSON(resp.Data)
	})

	products.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		resp, err := client.Service("products").GetByID(id)
		if err != nil {
			if err.Message == "project not found" {
				return c.Status(err.Status).JSON(map[string]interface{}{
					"message": "failed to connect to your project, please check if the api had been set properly.",
				})
			}

			return c.Status(err.Status).JSON(err)
		}

		return c.Status(http.StatusOK).JSON(resp.Data)
	})

	products.Patch("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		body := make(map[string]interface{})

		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(http.StatusBadRequest).SendString("failed parse request body to json")
		}

		resp, err := client.Service("products").UpdateByID(id, body)
		if err != nil {
			if err.Message == "project not found" {
				return c.Status(err.Status).JSON(map[string]interface{}{
					"message": "failed to connect to your project, please check if the api had been set properly.",
				})
			}

			return c.Status(err.Status).JSON(err)
		}

		return c.Status(http.StatusOK).JSON(resp.Data)
	})

	products.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		resp, err := client.Service("products").DeleteByID(id)
		if err != nil {
			if err.Message == "project not found" {
				return c.Status(err.Status).JSON(map[string]interface{}{
					"message": "failed to connect to your project, please check if the api had been set properly.",
				})
			}

			return c.Status(err.Status).JSON(err)
		}

		return c.Status(http.StatusOK).JSON(resp.Data)
	})

	app.Listen(":3000")
}
