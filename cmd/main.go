package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/rostekus/silvestrov/internal/models"
	"github.com/rostekus/silvestrov/internal/nats"
	"github.com/rostekus/silvestrov/internal/server/httpserver"
	"github.com/rostekus/silvestrov/internal/sqs"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	config := httpserver.SQSServerConfig{Port: 8080}

	nQS, err := queue.NewNATSQueueStorage("localhost:4222")
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot connect to NATS %v", err))
		os.Exit(-1)
	}

	q := models.QueueInfo{
		Name: "q",
	}
	q1 := models.QueueInfo{
		Name: "q1",
	}

	c := context.TODO()

	nQS.CreateQueue(c, 1, q)
	nQS.CreateQueue(c, 1, q1)

	fmt.Println(nQS.ListQueues(c, 1))

	nQS.Publish(c, 1, "q", []byte("name"))

	sqsClient := sqs.NewSQS(nQS)

	serverSQS := httpserver.NewServerSQS(logger, sqsClient, config)

	defer serverSQS.Stop()
	logger.Info(fmt.Sprintf("SQS Endpoint: http://localhost:%d\n", config.Port))
	if err := serverSQS.StartAndListen(); err != nil {

		logger.Error(fmt.Sprintf("Error %s, shutting down", err.Error()))
	}
	logger.Info("Shut down")

}
