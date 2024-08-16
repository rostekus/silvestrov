package sqs

import (
	"context"
	"fmt"
	"regexp"

	"github.com/rostekus/silvestrov/internal/models"
)

type SQS struct {
	queueStore models.QueueStorage
}

func NewSQS(qStore models.QueueStorage) *SQS {
	return &SQS{
		queueStore: qStore,
	}
}

func (s *SQS) GetQueue(tenantId int64, queueName string) (models.Queue, error) {

	return s.GetQueue(tenantId, queueName)

}
func (s *SQS) CreateQueue(c *context.Context, tenantId int64, qToStore models.Queue) (models.Queue, error) {
	if len(qToStore.Name) > 80 {
		return models.Queue{}, fmt.Errorf("Queue name error should be longer then 80 chars")
	}

	regex, err := regexp.Compile(`^[a-zA-Z0-9-_]+$`)
	if err != nil {
		return models.Queue{}, err
	}

	if !regex.MatchString(qToStore.Name) {
		return models.Queue{}, fmt.Errorf("Queue name is not in valid format")
	}

	qNew, err := s.queueStore.CreateQueue(tenantId, qToStore)

	if err != nil {
		return models.Queue{}, err
	}

	return qNew, err
}
