.PHONY: build
repo:
	docker-compose up -d --build

http:
	go run ./cmd/main.go

build: repo http


.DEFAULT_GOAL := build