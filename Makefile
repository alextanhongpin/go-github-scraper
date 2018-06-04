REPO := alextanhongpin/go-github-scraper

VERSION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -R)
VCS_URL := $(shell basename `git rev-parse --show-toplevel`)
VCS_REF := $(shell git log -1 --pretty=%h)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)

SEMVER_VERSION := 1.0.7

include .env

start:
	GITHUB_TOKEN=${GITHUB_TOKEN} DB_NAME=${DB_NAME} DB_HOST=${DB_HOST} DB_USER=${DB_USER} DB_PASS=${DB_PASS} go run main.go

mem:
	@go tool pprof --alloc_space http://localhost:6060/debug/pprof/heap

# Collect a 30-seconds cpu profiling
cpu:
	go tool pprof  http://localhost:6060:/debug/pprof/profile

docker:
	@docker build -t ${REPO} --build-arg VERSION="${VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg VCS_URL="${VCS_URL}" \
	--build-arg VCS_REF="${VCS_REF}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" .

tag: 
	@docker tag ${REPO}:latest ${REPO}:${SEMVER_VERSION}

push: 
	@docker push ${REPO}:latest
	@docker push ${REPO}:${SEMVER_VERSION}

deploy:
	make docker && make tag && make push