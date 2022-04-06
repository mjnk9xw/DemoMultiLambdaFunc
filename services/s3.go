package services

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grokify/go-awslambda"
)

func Upload(request events.APIGatewayProxyRequest, cfg aws.Config) (image string, err error) {
	client := s3.NewFromConfig(cfg)
	r, err := awslambda.NewReaderMultipart(request)
	if err != nil {
		return
	}

	part, err := r.NextPart()
	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(part)
	if err != nil {
		return
	}

	bucket := "serverless-series-upload-minhnv"
	filename := part.FileName()

	data := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   bytes.NewReader(content),
	}

	_, err = client.PutObject(context.TODO(), data)
	if err != nil {
		return
	}

	image = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, "ap-southeast-1", filename)

	return
}
