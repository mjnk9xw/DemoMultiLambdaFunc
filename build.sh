cd lambda/lambda_create
go build -o main main.go
zip create.zip main
rm -rf main