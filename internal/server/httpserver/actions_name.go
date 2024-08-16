package httpserver

type AWSMethod string

const (
	AmazonSQSSendMessage        AWSMethod = "AmazonSQS.SendMessage"
	AmazonSQSSendMessageBatch   AWSMethod = "AmazonSQS.SendMessageBatch"
	AmazonSQSReceiveMessage     AWSMethod = "AmazonSQS.ReceiveMessage"
	AmazonSQSDeleteMessage      AWSMethod = "AmazonSQS.DeleteMessage"
	AmazonSQSListQueues         AWSMethod = "AmazonSQS.ListQueues"
	AmazonSQSGetQueueUrl        AWSMethod = "AmazonSQS.GetQueueUrl"
	AmazonSQSCreateQueue        AWSMethod = "AmazonSQS.CreateQueue"
	AmazonSQSGetQueueAttributes AWSMethod = "AmazonSQS.GetQueueAttributes"
	AmazonSQSPurgeQueue         AWSMethod = "AmazonSQS.PurgeQueue"
)
