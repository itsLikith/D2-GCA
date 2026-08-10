package proxy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// Forward creates a middleware that proxies requests to
// the specified upstream service URL.
func Forward(
	upstreamURL string,
) fiber.Handler {

	return func(c *fiber.Ctx) error {

		url := upstreamURL + c.OriginalURL()

		if err := proxy.Do(c, url); err != nil {

			return c.Status(fiber.StatusBadGateway).JSON(
				fiber.Map{
					"code":    "PROXY_ERROR",
					"message": err.Error(),
				},
			)
		}

		return nil
	}
}
