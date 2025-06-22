.PHONY: all run build clean

all: deps run

deps:
	@go mod tidy

run: build
	@./bin/shodan-cli

build:
	@go build -o bin/shodan-cli

clean:
	@rm -rf bin
