package delivery

import (
	"Study/AWS/DemoMultiLambdaFunc/entities"
	"Study/AWS/DemoMultiLambdaFunc/services"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
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

	image, err := services.Upload(req, cfg)
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
	res, err := services.Create(&entities.Book{
		ID:     form.Value["id"][0],
		Name:   form.Value["name"][0],
		Author: form.Value["author"][0],
		Image:  image,
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

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(res),
	}, nil
}
