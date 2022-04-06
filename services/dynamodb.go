package services

import (
	"Study/AWS/DemoMultiLambdaFunc/entities"
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Create(req *entities.Book) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	svc := dynamodb.NewFromConfig(cfg)
	data, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("books"),
		Item: map[string]types.AttributeValue{
			"id":     &types.AttributeValueMemberS{Value: req.ID},
			"name":   &types.AttributeValueMemberS{Value: req.Name},
			"author": &types.AttributeValueMemberS{Value: req.Author},
			"image":  &types.AttributeValueMemberS{Value: req.Image},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return "", err
	}

	res, _ := json.Marshal(data)
	return string(res), nil
}
