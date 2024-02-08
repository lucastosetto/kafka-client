.PHONY: build
build:
	@if ! command -v go &> /dev/null; then \
		echo "\nGo is not installed. Please install Go before proceeding."; \
		echo "Visit https://golang.org/doc/install for installation instructions.\n"; \
		exit 1; \
	fi
	@echo "\nBuilding kafka-client binary...\n"
	@go build -o kafka-client ./src/main.go
	@echo "kafka-client binary built successfully!\n"
	@echo "To use kafka-client from anywhere on your system, add it to your PATH."
	@echo "Example:"
	@echo "    export PATH=\$$PATH:${PWD}\n"
