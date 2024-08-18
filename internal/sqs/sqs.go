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

func (s *SQS) PublishMsg(c context.Context, tenantId int64, queueName string, msg []byte) error {
	return s.queueStore.Publish(c, tenantId, queueName, msg)
}

func (s *SQS) GetQueue(c context.Context, tenantId int64, queueName string) (models.QueueInfo, error) {

	return s.queueStore.GetQueue(c, tenantId, queueName)

}
func (s *SQS) CreateQueue(c context.Context, tenantId int64, qToStore models.QueueInfo) (models.QueueInfo, error) {
	if len(qToStore.Name) > 80 {
		return models.QueueInfo{}, fmt.Errorf("Queue name error should be longer then 80 chars")
	}

	regex, err := regexp.Compile(`^[a-zA-Z0-9-_]+$`)
	if err != nil {
		return models.QueueInfo{}, err
	}

	if !regex.MatchString(qToStore.Name) {
		return models.QueueInfo{}, fmt.Errorf("Queue name is not in valid format")
	}

	qNew, err := s.queueStore.CreateQueue(c, tenantId, qToStore)

	if err != nil {
		return models.QueueInfo{}, err
	}

	return qNew, err
}
