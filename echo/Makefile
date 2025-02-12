BIN       = bin/echo-test

.PHONY: usage
usage:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build  - Build the binary"
	@echo "  run    - Run development server"
	@echo "  clean  - Remove the binary"
	@echo "  lint   - Run go fmt"

.PHONY: lint
lint:
	go fmt ./...

.PHONY: builld
build:
	go build -o $(BIN) cmd/echo-test/main.go

.PHONY: run
run:
	go run cmd/echo-test/main.go

.PHONY: clean
clean:
	rm -rf bin
