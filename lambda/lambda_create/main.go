package main

import (
	"Study/AWS/DemoMultiLambdaFunc/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(services.Create)
}
