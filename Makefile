.PHONY: init build run

all: init build run

init:
	go mod tidy
	go mod vendor

build:
	go build -o build/single-fizz-buzz.exe cmd/main.go 

run:
	./build/single-fizz-buzz.exe