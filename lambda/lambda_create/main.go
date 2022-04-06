package main

import (
	"Study/AWS/DemoMultiLambdaFunc/delivery"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(delivery.Create)
}
