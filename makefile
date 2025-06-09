compile: build run

build: 
	go build -o bin/compiler cmd/main.go

run: 
	./bin/compiler