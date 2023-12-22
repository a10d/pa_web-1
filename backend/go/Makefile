
dev:
	@ APP_INTERFACE=localhost:8080 air

build:
	@go build -o bin/backend

run: build
	@./bin/backend

test:
	@go test -v ./...
