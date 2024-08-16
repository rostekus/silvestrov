package models

type Queue struct {
	id                int
	Name              string
	RateLimit         float64
	MaxRetries        int
	VisibilityTimeout int
}

type QueueStorage interface {
	GetQueue(tenantId int64, queueName string) (Queue, error)
	CreateQueue(tenantId int64, queue Queue) (Queue, error)
}
