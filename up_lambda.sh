cd lambda

echo "up lambda create"
cd lambda_create
aws lambda update-function-code --function-name books_create --zip-file fileb://create.zip --region ap-southeast-1

echo "up lambda get list"
cd ../lambda_getlist
aws lambda update-function-code --function-name books_list --zip-file fileb://list.zip --region ap-southeast-1
