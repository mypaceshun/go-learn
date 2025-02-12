BIN       = bin/echo-test

.PHONY: usage
usage:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build  - Build the binary"
	@echo "  run    - Build and run the binary"
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
	${MAKE} build
	$(BIN)

.PHONY: clean
clean:
	rm -rf bin
