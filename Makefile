.PHONY: build
build:
	rm -rf build 
	mkdir build
	go build -o build/server -v ./cmd

.PHONY: run
run:
	go run cmd/main.go

.DEFAULT_GOAL := build