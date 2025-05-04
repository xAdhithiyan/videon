build: 
	@go build -o bin/output cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/output

serve: 
	@air --build.cmd "go build -o bin/output cmd/main.go" --build.bin "./bin/output"