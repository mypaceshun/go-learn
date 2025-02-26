.PHONY: usage
usage:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run: run the project"
	@echo "  test: run tests"
	@echo "  format: format the code"


.PHONY: run
run:
	go run main.go

.PHONY: test
test: format
	go test -v ./...

.PHONY: format
format:
	go fmt ./...
