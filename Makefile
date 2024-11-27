build:
	@go build -o bin/compression_tool cmd/main.go

run: build
	@./bin/compression_tool

test:
	@go test ./... -v
