cd lambda

echo "build lambda create"
cd lambda_create
go build -o main main.go
zip create.zip main 
rm -rf main

echo "build lambda get list"
cd ../lambda_getlist
go build -o main main.go
zip list.zip main
rm -rf main