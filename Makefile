include .env

start:
	GITHUB_TOKEN=${GITHUB_TOKEN} DB_NAME=${DB_NAME} DB_HOST=${DB_HOST} DB_USER=${DB_USER} DB_PASS=${DB_PASS} go run main.go

alloc:
	@go tool pprof -alloc_space -svg http://localhost:8080/debug/pprof/heap > heap.svg

heap:
	@go tool pprof -png http://localhost:8080/debug/pprof/heap > out.png
