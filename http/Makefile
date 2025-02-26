.PHONY: usage
usage:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build: build the project"
	@echo "  clean: clean the project"
	@echo "  run: run the project"
	@echo "  test: run the tests"
	@echo "  usage: print this message"

.PHONY: build
build:
	@echo "Building the project..."
	@make bin/client
	@make bin/server

bin:
	@mkdir -p bin

bin/client: bin
	@echo "Building client..."
	@go build -o bin/client ./cmd/client
	@chmod +x bin/client

bin/server: bin
	@echo "Building server..."
	@go build -o bin/server ./cmd/server
	@chmod +x bin/server

.PHONY: clean
clean:
	@echo "Cleaning the project..."
	@rm -rf bin

.PHONY: run
run:
	@echo "Running the project..."
	@make build
	@./bin/server

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
