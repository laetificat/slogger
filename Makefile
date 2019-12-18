.DEFAULT_GOAL := help

help:
	@echo "Available options:\n- test: Run all tests\n- clean: Clean project\n- check: Test and check project"

test:
	@echo "Running gotest..."
	@gotest ./... -coverprofile=coverage.out -count=1

clean:
	@echo "Cleaning project..."
	@rm -f coverage.out

check: test
	@golangci-lint run