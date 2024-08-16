package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SlogMiddleware(logger *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process the request
		err := c.Next()

		// Log the request details after processing
		logger.Info("Request",
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.Int("status", c.Response().StatusCode()),
			slog.Duration("duration", time.Since(start)),
		)

		return err
	}
}
