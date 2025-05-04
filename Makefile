build: 
	@go build -o bin/output cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/output
