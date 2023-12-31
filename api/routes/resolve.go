package routes

import (
	"github.com/ayushgupta4002/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connectivity issue",
		})
	}
	rinr := database.CreateClient(1)
	defer rinr.Close()
	_ = rinr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}
