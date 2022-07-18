.PHONY: test t build run image ctn-run publish

CTN ?= podman
REGISTRY ?= docker.io/albandewilde

test t:
	@go test ./...

build:
	@go build -o ./intech-bot

run: build
	@TKN=$(TKN) ./intech-bot

image:
	@$(CTN) build -t $(REGISTRY)/intech-bot .

ctn-run: image
	@$(CTN) run \
		-e TKN=$(TKN) \
		intech-bot

publish:
	@$(CTN) push $(REGISTRY)/intech-bot
