build:
	go build -o bin/navi cmd/navi/main.go

run:
	go run cmd/navi/main.go

install:
	cd cmd/navi && go install

format:
	go fmt ./...
