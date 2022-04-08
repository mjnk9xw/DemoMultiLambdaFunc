echo "build lambda get list"
cd lambda/lambda_getlist
go build -o main main.go
zip list.zip main
rm -rf main
