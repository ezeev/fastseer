package storage

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/ezeev/fastseer/shopify"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestDynamoOpen(t *testing.T) {

	db := DynamoDbStorage{}

	options := make(map[string]string)
	options["aws_zone"] = "us-west-2"
	options["creds"] = "../secret/credentials"
	options["table"] = "clients"
	options["keyField"] = "shop"

	err := db.Open(options)
	if err != nil {
		t.Error(err)
	}
	// can we read clients table?
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("clients"),
	}
	result, err := db.db.DescribeTable(req)
	if err != nil {
		t.Log(err)
	}
	t.Log(result.GoString())
}

func TestDynamoImplPutAndGet(t *testing.T) {

	options := make(map[string]string)
	options["aws_zone"] = "us-west-2"
	options["creds"] = "../secret/credentials"
	options["table"] = "clients"
	options["keyField"] = "shop"

	db, _ := NewStorage("dynamo")

	err := db.Open(options)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	client := shopify.ShopifyClientConfig{}
	client.Shop = "acme"
	client.IndexAddress = "http://solrurl.com"

	err = db.Put(client.Shop, client)
	if err != nil {
		t.Error(err)
	}

	var newClient shopify.ShopifyClientConfig
	err = db.Get("acme", &newClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(newClient)

}
