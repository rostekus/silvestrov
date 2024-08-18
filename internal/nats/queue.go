package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rostekus/silvestrov/internal/models"
)

type NATSQueueStorage struct {
	conn *nats.Conn
	kv   nats.KeyValue
}

func NewNATSQueueStorage(natsURL string) (*NATSQueueStorage, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	kv, err := js.KeyValue("queues")
	if err != nil {
		// If the KV store doesn't exist, create it
		kv, err = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: "queues",
		})
		if err != nil {
			return nil, err
		}
	}

	return &NATSQueueStorage{
		conn: conn,
		kv:   kv,
	}, nil
}

func (n *NATSQueueStorage) Publish(c context.Context, tenantId int64, queueName string, message []byte) error {
	_, err := n.GetQueue(c, tenantId, queueName)
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("tenant_%d.queue.%s", tenantId, queueName)
	print(subject)

	return n.conn.Publish(subject, message)
}

func (n *NATSQueueStorage) GetQueue(c context.Context, tenantId int64, queueName string) (models.QueueInfo, error) {
	key := fmt.Sprintf("%d_%s", tenantId, queueName)
	entry, err := n.kv.Get(key)
	if err != nil {
		if errors.Is(err, nats.ErrKeyNotFound) {
			return models.QueueInfo{}, errors.New("queue not found")
		}
		return models.QueueInfo{}, err
	}
	var queueInfo models.QueueInfo
	err = json.Unmarshal(entry.Value(), &queueInfo)
	if err != nil {
		return models.QueueInfo{}, err
	}

	return queueInfo, nil
}

func (n *NATSQueueStorage) CreateQueue(c context.Context, tenantId int64, queue models.QueueInfo) (models.QueueInfo, error) {
	key := fmt.Sprintf("%d_%s", tenantId, queue.Name)
	data, err := json.Marshal(queue)
	if err != nil {
		return models.QueueInfo{}, err
	}

	_, err = n.kv.Put(key, data)
	if err != nil {
		return models.QueueInfo{}, err
	}

	return queue, nil
}

func (n *NATSQueueStorage) DeleteQueue(c context.Context, tenantId int64, queueName string) error {
	key := fmt.Sprintf("%d_%s", tenantId, queueName)
	err := n.kv.Delete(key)
	if err != nil {
		if errors.Is(err, nats.ErrKeyNotFound) {
			return errors.New("queue not found")
		}
		return err
	}

	return nil
}

func (n *NATSQueueStorage) ListQueues(c context.Context, tenantId int64) ([]string, error) {
	keys, err := n.kv.Keys()
	if err != nil {
		return nil, err
	}

	var queues []string
	for _, key := range keys {
		if len(key) > 0 && key[:len(fmt.Sprintf("%d_", tenantId))] == fmt.Sprintf("%d_", tenantId) {
			queues = append(queues, key[len(fmt.Sprintf("%d_", tenantId)):])
		}
	}

	return queues, nil
}
