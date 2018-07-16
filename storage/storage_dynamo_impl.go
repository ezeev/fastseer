package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ezeev/fastseer/util"
)

/*
type Storage interface {
	Open() error
	Get(key string, val *interface{}) error
	Put(key string, val interface{}) error
	Close() error
}

*/

type DynamoDbStorage struct {
	sess *session.Session
	db   *dynamodb.DynamoDB
	opts map[string]string
}

func (d *DynamoDbStorage) Open(options map[string]string) error {

	d.opts = options

	err := util.ValidateOptionsMap(options, "aws_zone", "creds", "table", "keyField")
	if err != nil {
		return err
	}

	zone := d.opts["aws_zone"]   // us-west-2
	credsPath := d.opts["creds"] // "secret/credentials"

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(zone),
		Credentials: credentials.NewSharedCredentials(credsPath, ""),
	})
	d.sess = sess
	d.db = dynamodb.New(sess)
	return err
}

func (d *DynamoDbStorage) Get(key string, val interface{}) error {

	table := d.opts["table"]       // clients
	keyField := d.opts["keyField"] // shop

	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			keyField: {
				S: aws.String(key),
			},
		},
	}

	result, err := d.db.GetItem(input)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, val)
	return err
}

func (d *DynamoDbStorage) Put(key string, val interface{}) error {

	// Key isn't required for DynamoDB, instead, you must have an
	// attribute specified as the Id when the table is created
	av, err := dynamodbattribute.MarshalMap(val)
	if err != nil {
		return err
	}

	table := d.opts["table"]

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = d.db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDbStorage) Close() error {
	return nil
}
