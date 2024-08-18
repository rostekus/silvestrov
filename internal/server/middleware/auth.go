package middleware

import (
	"github.com/gofiber/fiber/v2"
)

const TenantID = "tenantID"

func AuthMiddleware(c *fiber.Ctx) error {
	c.Locals(TenantID, int64(1))
	return c.Next()
}
