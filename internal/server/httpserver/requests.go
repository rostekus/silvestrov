package httpserver

type CreateQueueRequest struct {
	QueueName  string            `json:"QueueName"`
	Attributes map[string]string `json:"Attributes,omitempty"`
	Tags       map[string]string `json:"Tags,omitempty"`
}

type PublishMessageRequest struct {
	QueueUrl               string                           `json:"QueueUrl"`
	MessageBody            string                           `json:"MessageBody"`
	DelaySeconds           int                              `json:"DelaySeconds,omitempty"`
	MessageAttributes      map[string]MessageAttributeValue `json:"MessageAttributes,omitempty"`
	MessageDeduplicationId string                           `json:"MessageDeduplicationId,omitempty"`
	MessageGroupId         string                           `json:"MessageGroupId,omitempty"`
}

type MessageAttributeValue struct {
	StringValue     string   `json:"StringValue,omitempty"`
	BinaryValue     string   `json:"BinaryValue,omitempty"`
	StringListValue []string `json:"StringListValue,omitempty"`
	BinaryListValue [][]byte `json:"BinaryListValue,omitempty"`
	DataType        string   `json:"DataType"`
}
