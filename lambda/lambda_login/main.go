// package main

// import (
// 	"Study/AWS/DemoMultiLambdaFunc/delivery"

// 	"github.com/aws/aws-lambda-go/lambda"
// )

// func main() {
// 	lambda.Start(delivery.Login)
// }

package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body Body
	log.Println("[login] body = ", req.Body)
	b64String, _ := base64.StdEncoding.DecodeString(req.Body)
	rawIn := json.RawMessage(b64String)
	bodyBytes, err := rawIn.MarshalJSON()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while retrieving AWS credentials",
		}, nil
	}

	cip := cognitoidentityprovider.NewFromConfig(cfg)
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String("6bc1ep4mpjjl4upv37kgtbu8qa"), // Should os.Getenv("CLIENT_ID")
		AuthParameters: map[string]string{
			"USERNAME": body.Username,
			"PASSWORD": body.Password,
			// "USERNAME": "minhnv98vp@gmail.com",
			// "PASSWORD": "@#Minh123@#",
		},
	}
	authResp, err := cip.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	res, _ := json.Marshal(authResp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(res),
	}, nil
}

func main() {
	lambda.Start(login)
}
