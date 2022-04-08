cd lambda

echo "up lambda create"
cd lambda_create
aws lambda update-function-code --function-name books_create --zip-file fileb://create.zip --region ap-southeast-1

echo "up lambda get list"
cd ../lambda_getlist
aws lambda update-function-code --function-name books_list --zip-file fileb://list.zip --region ap-southeast-1


echo "up lambda login"
cd ../lambda_login
aws lambda update-function-code --function-name login --zip-file fileb://login.zip --region ap-southeast-1


echo "up lambda change pass"
cd ../lambda_changepassword
aws lambda update-function-code --function-name change_password --zip-file fileb://change-password.zip --region ap-southeast-1
