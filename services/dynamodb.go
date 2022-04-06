package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/grokify/go-awslambda"
)

func Create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: "Error while retrieving AWS credentials",
		}, nil
	}

	image, err := Upload(req, cfg)
	// image := "https://serverless-series-upload-minhnv.s3.ap-southeast-1.amazonaws.com/2022-01-28-103246.jpg"
	// err = nil
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	r, _ := awslambda.NewReaderMultipart(req)
	form, _ := r.ReadForm(1024)
	svc := dynamodb.NewFromConfig(cfg)
	data, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("books"),
		Item: map[string]types.AttributeValue{
			"id":     &types.AttributeValueMemberS{Value: form.Value["id"][0]},
			"name":   &types.AttributeValueMemberS{Value: form.Value["name"][0]},
			"author": &types.AttributeValueMemberS{Value: form.Value["author"][0]},
			"image":  &types.AttributeValueMemberS{Value: image},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	res, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(res),
	}, nil
}
