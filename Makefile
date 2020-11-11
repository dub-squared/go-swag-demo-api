.PHONY: build
## build: Compile the packages.
build:
	@go build -o ./cmd/rest-api ./cmd/rest-api

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./cmd/rest-api/rest-api -p 3001