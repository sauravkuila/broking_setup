MOCKGEN := $(shell command -v mockgen 2> /dev/null)

build:
	go build -o broking_setup cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v -cover ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm -f coverage.out