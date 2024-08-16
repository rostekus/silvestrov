package httpserver

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/rostekus/silvestrov/internal/models"
	"github.com/rostekus/silvestrov/internal/server/middleware"
)

func (s *sqsServer) CreateQueue(c *fiber.Ctx) error {

	req := &CreateQueueRequest{}

	err := json.Unmarshal(c.Body(), req)
	if err != nil {
		return err
	}

	tenantId := c.Locals(middleware.TenantID).(int64)

	q := models.Queue{
		Name:              req.QueueName,
		RateLimit:         -1,
		MaxRetries:        -1,
		VisibilityTimeout: 30,
	}

	ctx := context.Background()

	_, err = s.sqs.CreateQueue(&ctx, tenantId, q)

	return err

}
