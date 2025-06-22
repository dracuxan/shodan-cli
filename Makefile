.PHONY: all run build clean

deps:
	@go mod tidy

run: build
	@./bin/shodan-cli

build:
	@go build -o bin/shodan-cli

clean:
	@rm -rf bin
