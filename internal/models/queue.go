package models

import "context"

type QueueInfo struct {
	id                int
	Name              string
	RateLimit         float64
	MaxRetries        int
	VisibilityTimeout int
}

type QueueStorage interface {
	GetQueue(c context.Context, tenantId int64, queueName string) (QueueInfo, error)
	CreateQueue(c context.Context, tenantId int64, queue QueueInfo) (QueueInfo, error)
	DeleteQueue(c context.Context, tenantId int64, queueName string) error
	ListQueues(c context.Context, tenantId int64) ([]string, error)
	Publish(c context.Context, tenantId int64, queueName string, message []byte) error
}
