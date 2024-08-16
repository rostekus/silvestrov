package middleware

import (
	"github.com/gofiber/fiber/v2"
)

const TenantID = "tenantID"

func AuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
