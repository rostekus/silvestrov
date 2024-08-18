package httpserver

type CreateQueueResponse struct {
	QueueUrl string `json:"QueueUrl"`
}

type PublishMessageResponse struct {
	MD5OfMessageBody       string `json:"MD5OfMessageBody"`
	MD5OfMessageAttributes string `json:"MD5OfMessageAttributes"`
	MessageId              string `json:"MessageId"`
	SequenceNumber         string `json:"SequenceNumber,omitempty"`
}
