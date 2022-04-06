cd lambda/lambda_create
aws lambda update-function-code --function-name books_create --zip-file fileb://create.zip --region ap-southeast-1