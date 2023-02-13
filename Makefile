build:
	@go build -o bin/bbank

run: build
	@./bin/bbank

test:
	@go test -v ./...