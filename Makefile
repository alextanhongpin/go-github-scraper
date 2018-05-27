include .env

start:
	GITHUB_TOKEN=${GITHUB_TOKEN} DB_NAME=${DB_NAME} DB_HOST=${DB_HOST} go run main.go

