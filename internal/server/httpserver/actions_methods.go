package httpserver

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/rostekus/silvestrov/internal/models"
	"github.com/rostekus/silvestrov/internal/server/middleware"
)

func (s *sqsServer) PublishMsg(c *fiber.Ctx) error {

	req := &PublishMessageRequest{}

	err := json.Unmarshal(c.Body(), req)
	if err != nil {
		return err
	}

	tenantId := s.getTenantID(c)

	ctx := context.Background()

	err = s.sqs.PublishMsg(ctx, tenantId, req.QueueUrl, []byte(req.MessageBody))
	if err != nil {
		return err
	}

	resp := PublishMessageResponse{
		MessageId: "1",
	}

	return c.JSON(resp)

}

func (s *sqsServer) CreateQueue(c *fiber.Ctx) error {

	req := &CreateQueueRequest{}

	err := json.Unmarshal(c.Body(), req)
	if err != nil {
		return err
	}

	tenantId := s.getTenantID(c)

	q := models.QueueInfo{
		Name:              req.QueueName,
		RateLimit:         -1,
		MaxRetries:        -1,
		VisibilityTimeout: 30,
	}

	ctx := context.Background()

	qCreated, err := s.sqs.CreateQueue(ctx, tenantId, q)

	if err != nil {
		return err
	}
	resp := CreateQueueResponse{
		QueueUrl: qCreated.Name,
	}
	return c.JSON(resp)

}

func (s *sqsServer) getTenantID(c *fiber.Ctx) int64 {

	return c.Locals(middleware.TenantID).(int64)
}
