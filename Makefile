build: 
	@go build -o bin/output cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/output

serve: 
	@air --build.cmd "go build -o bin/output cmd/main.go" --build.bin "./bin/output"

migration:
	@migrate create -ext sql -dir db/migration $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migration/migrate.go up

migrate-down:
	@go run cmd/migration/migrate.go down