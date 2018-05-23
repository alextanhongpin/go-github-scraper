include .env

start:
	@echo ${GITHUB_TOKEN}
	@go run main.go -github_token=${GITHUB_TOKEN}

