package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayushgupta4002/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", mainpage)
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenURL)

}
func mainpage(c *fiber.Ctx) error {
	err := c.JSON(fiber.Map{
		"message": "server up and kicking",
	})
	if err != nil {
		// Handle the error if needed
		return err
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("hi")
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
