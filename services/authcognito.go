package services

import (
	"Study/AWS/DemoMultiLambdaFunc/data/body"
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func ChangePassword(bodydata body.BodyChangePass) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	cip := cognitoidentityprovider.NewFromConfig(cfg)
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String("6bc1ep4mpjjl4upv37kgtbu8qa"),
		AuthParameters: map[string]string{
			"USERNAME": bodydata.Username,
			"PASSWORD": bodydata.OldPassword,
		},
	}
	authResp, err := cip.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		return "", err
	}

	challengeInput := &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName: "NEW_PASSWORD_REQUIRED",
		ClientId:      aws.String("6bc1ep4mpjjl4upv37kgtbu8qa"),
		ChallengeResponses: map[string]string{
			"USERNAME":     bodydata.Username,
			"NEW_PASSWORD": bodydata.NewPassword,
		},
		Session: authResp.Session,
	}
	challengeResp, err := cip.RespondToAuthChallenge(context.TODO(), challengeInput)
	if err != nil {
		return "", err
	}

	res, _ := json.Marshal(challengeResp)
	return string(res), nil
}

func Login(bodydata body.BodyLogin) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	cip := cognitoidentityprovider.NewFromConfig(cfg)
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String("6bc1ep4mpjjl4upv37kgtbu8qa"),
		AuthParameters: map[string]string{
			"USERNAME": bodydata.Username,
			"PASSWORD": bodydata.Password,
		},
	}
	authResp, err := cip.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		return "", err
	}

	res, _ := json.Marshal(authResp)
	return string(res), nil
}
