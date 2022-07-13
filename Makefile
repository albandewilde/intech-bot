.PHONY: test t build run ctn-build ctn-run

CTN ?= podman

test t:
	@go test ./...

build:
	@go build -o ./intech-bot

run: build
	@./intech-bot

ctn-build:
	@$(CTN) build -t intech-bot .

ctn-run: ctn-build
	@$(CTN) run -t intech-bot
