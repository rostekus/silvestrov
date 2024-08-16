package httpserver

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/rostekus/silvestrov/internal/server/middleware"
	"github.com/rostekus/silvestrov/internal/sqs"
)

const methodHeader = "X-Amz-Target"

type SQSServerConfig struct {
	Port int
}

type sqsServer struct {
	logger *slog.Logger
	sqs    *sqs.SQS
	server *fiber.App
	cfg    SQSServerConfig
}

func NewServerSQS(logger *slog.Logger, sqsClient *sqs.SQS, config SQSServerConfig) *sqsServer {
	app := fiber.New()

	app.Use(middleware.SlogMiddleware(logger))

	s := &sqsServer{logger: logger,
		sqs:    sqsClient,
		server: app,
		cfg:    config,
	}
	app.Post("/*", s.handleRequest)
	return s
}

func (s *sqsServer) StartAndListen() error {

	return s.server.Listen(fmt.Sprintf(":%d", s.cfg.Port))
}

func (s *sqsServer) Stop() error {
	return s.server.Shutdown()

}

// TODO: Handle actions in goroutine
func (s *sqsServer) handleRequest(c *fiber.Ctx) error {

	action, err := s.getAction(c)

	if err != nil {
		return err
	}
	handleMsg := "Handling request action %s"

	awsMethod := AWSMethod(action)
	s.logger.Info(fmt.Sprintf(handleMsg, awsMethod))
	var rc error
	switch awsMethod {
	case AmazonSQSSendMessage:
	case AmazonSQSSendMessageBatch:
	case AmazonSQSReceiveMessage:
	case AmazonSQSDeleteMessage:
	case AmazonSQSListQueues:
	case AmazonSQSGetQueueUrl:
	case AmazonSQSCreateQueue:
		rc = s.CreateQueue(c)
	case AmazonSQSGetQueueAttributes:
	case AmazonSQSPurgeQueue:
	default:
		rc = fmt.Errorf("Undefined action")
	}

	return rc
}

func (s *sqsServer) getAction(c *fiber.Ctx) (string, error) {

	awsMethodHeader, ok := c.GetReqHeaders()[methodHeader]
	if !ok {
		errMsg := fmt.Sprintf("%s header not found", methodHeader)
		return "", fmt.Errorf(errMsg)
	}
	methodStr := awsMethodHeader[0]

	return methodStr, nil

}
