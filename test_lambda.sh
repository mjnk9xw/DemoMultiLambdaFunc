echo "test create book"

aws lambda invoke --function-name books_create response.json

echo "test get list book"

aws lambda invoke --function-name books_list response.json

echo "test login"

aws lambda invoke --function-name login response.json

echo "test change-password"

aws lambda invoke --function-name change_password response.json