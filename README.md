# GO-PLUGIN

Developed an Plugin Integration for Collect Systems Architecture which can be used for different actions through generating an `.so` file.
In this project, we are developing a system for data collection for forms which will include user details(identification through id), forms detail, questions, responses detail and answers.

# Design Doc 

https://docs.google.com/document/d/1SJWn7xr43sQRojPts8QYsd40FdL5xhjitqP7vcYp1JI/edit?usp=sharing

# Setting up

Follow the following steps:

1. Create your `.env` file and add you mongoDB URI in the variable name `MONGODB_URI`
2. You can add your port using `PORT` env variable. By default it is 3000.
3. Type `go get` or `go mod tidy` to fetch all the packages required.
4. To generate `.so file` run this : `go build -buildmode=plugin -o actions/<ACTION_NAME>/action.so actions/<ACTION_NAME>/action.go`, ACTION_NAME = Google-Sheets for our use case.
5. Type `go run main.go` to start server.
6. You can use Postman to perform request.
7. You'll get the spreadsheet link in the terminal.