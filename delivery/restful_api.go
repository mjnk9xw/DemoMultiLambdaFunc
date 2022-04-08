package delivery

import (
	"Study/AWS/DemoMultiLambdaFunc/data/body"
	"Study/AWS/DemoMultiLambdaFunc/entities"
	"Study/AWS/DemoMultiLambdaFunc/services"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
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

func List(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res, err := services.GetList()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*", // add here
			},
			Body: "",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*", // add here
		},
		Body: res,
	}, nil
}

func ChangePassword(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var bodydata body.BodyChangePass
	b64String, _ := base64.StdEncoding.DecodeString(req.Body)
	rawIn := json.RawMessage(b64String)
	bodyBytes, err := rawIn.MarshalJSON()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	err = json.Unmarshal(bodyBytes, &bodydata)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	res, err := services.ChangePassword(bodydata)
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
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: res,
	}, nil
}

func Login(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var bodydata body.BodyLogin
	b64String, err := base64.StdEncoding.DecodeString(req.Body)
	log.Println("[Login] b64String = ", string(b64String), err)

	rawIn := json.RawMessage(b64String)
	bodyBytes, err := rawIn.MarshalJSON()

	log.Println("[Login] body = ", string(bodyBytes))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	err = json.Unmarshal(bodyBytes, &bodydata)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: err.Error(),
		}, nil
	}

	res, err := services.Login(bodydata)
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
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: res,
	}, nil
}
