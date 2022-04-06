echo "test create book"

aws lambda invoke --function-name books_create response.json

echo "test get list book"

aws lambda invoke --function-name books_list response.json