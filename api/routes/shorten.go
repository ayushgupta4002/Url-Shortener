package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ayushgupta4002/helpers"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	Url         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	//rate limiting

	//checking url
	if !govalidator.IsURL(body.Url) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Url",
		})
	}

	//domain restriction
	if !helpers.RemovedomainError(body.Url) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Invalid URL - bad domain error",
		})
	}

	//enforcing http:// before url
	body.Url = helpers.EnforceHttp(body.Url)

}
