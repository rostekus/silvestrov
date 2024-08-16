package httpserver

type CreateQueueRequest struct {
	QueueName  string            `json:"QueueName"`
	Attributes map[string]string `json:"Attributes,omitempty"`
	Tags       map[string]string `json:"Tags,omitempty"`
}
