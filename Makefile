build:
	go build -o bin/navigate cmd/navigate/main.go

run:
	go run cmd/navigate/main.go

install:
	cd cmd/navigate && go install

format:
	go fmt ./...
