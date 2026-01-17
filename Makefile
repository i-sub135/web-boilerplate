.PHONY: dev build run clean help

help:
	@echo "Available commands:"
	@echo "  make dev    - Run with hot reload using entr"
	@echo "  make build  - Build binary to ./tmp/main"
	@echo "  make run    - Run without hot reload"
	@echo "  make clean  - Clean tmp directory"

dev:
	find . \( -name "*.go" -o -name "*.html" \) | entr -r go run .

build:
	mkdir -p tmp
	go build -o ./tmp/main .

run:
	go run .

clean:
	rm -rf tmp/
