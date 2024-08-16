package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rostekus/silvestrov/internal/models"
	"github.com/rostekus/silvestrov/internal/server/httpserver"
	"github.com/rostekus/silvestrov/internal/sqs"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	config := httpserver.SQSServerConfig{Port: 8080}

	sqsClient := sqs.NewSQS(&FakeQueueStorage{})

	serverSQS := httpserver.NewServerSQS(logger, sqsClient, config)

	defer serverSQS.Stop()
	logger.Info(fmt.Sprintf("SQS Endpoint: http://localhost:%d\n", config.Port))
	if err := serverSQS.StartAndListen(); err != nil {

		logger.Error(fmt.Sprintf("Error %s, shutting down", err.Error()))
	}
	logger.Info("Shutted down")

}

type FakeQueueStorage struct {
}

func (s *FakeQueueStorage) GetQueue(tenantId int64, queueName string) (models.Queue, error) {
	return models.Queue{}, nil
}

// CreateQueue creates a new queue for a given tenantId
func (s *FakeQueueStorage) CreateQueue(tenantId int64, queue models.Queue) (models.Queue, error) {
	return models.Queue{}, nil
}
