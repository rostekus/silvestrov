import boto3


sqs = boto3.client(
    'sqs',
    aws_access_key_id='your-access-key-id',
    aws_secret_access_key='your-secret-access-key',
    region_name='your-region',
    endpoint_url="http://localhost:8080"
)
queue_name = 'MyTestQueue'
response = sqs.create_queue(QueueName=queue_name)
queue_url = response['QueueUrl']

print(f"Queue created successfully: {queue_url}")

message_body = 'Hello, this is a test message!'
response = sqs.send_message(
    QueueUrl=queue_url,
    MessageBody=message_body
)

print(f"Message sent successfully! Message ID: {response['MessageId']}")

