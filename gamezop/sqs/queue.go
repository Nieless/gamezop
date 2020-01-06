package gamezopSqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"strconv"
)

// Client is sqs client
type Client struct {
	QClient *sqs.SQS
	QUrl    *string
}

// GameStore defines the operations of sqs store.
type SqsStore interface {
	ReceiveMessages()(map[string]*string, error)
	DeleteMsg(string2 *string)
}

// NewClient returns a new sqs client.
func NewClient(config *Config) (*Client, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	}))

	client := &Client{
		QClient: sqs.New(sess),
	}
	return client, nil
}

// return a list of messages from sqs
func (client *Client) ReceiveMessages() (map[string]*string, error) {

	messageAttributeNames := []*string{aws.String(sqs.QueueAttributeNameAll)}
	attributeNames := []*string{aws.String(sqs.MessageSystemAttributeNameSentTimestamp)}

	receiveMessageInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: messageAttributeNames,
		AttributeNames:        attributeNames,
		QueueUrl:              client.QUrl,
		MaxNumberOfMessages:   aws.Int64(10),
		VisibilityTimeout:     aws.Int64(300),
		WaitTimeSeconds:       aws.Int64(10),
	}

	//data := make([]*string, 0)
	data := make(map[string]*string,0)

	result, err := client.QClient.ReceiveMessage(receiveMessageInput)
	if err != nil {
		return data, err
	}

	if result == nil || len(result.Messages) == 0 {
		log.Printf("nothing received from sqs: result is nil")
		return data, nil
	}

	for i, msg := range result.Messages {
		if msg.ReceiptHandle != nil{
			data[*msg.ReceiptHandle] = msg.Body
			continue
		}
		data[strconv.Itoa(i)] = msg.Body
	}

	return data, nil
}

// DeleteMsg deletes a new sqs client.
func (client Client) DeleteMsg(receiptHandle *string) {
	_, err := client.QClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      client.QUrl,
		ReceiptHandle: receiptHandle,
	})

	if err != nil {
		log.Printf("error while deleting msg from sqs: %s\n", err.Error())
	}
}
