.PHONY: test t build run image ctn-run publish

CTN ?= podman
REGISTRY ?= docker.io/albandewilde
HOST ?= 0.0.0.0
PORT ?= 5419

test t:
	@go test ./...

build:
	@go build -o ./intech-bot

run: build
	@TKN=$(TKN) HOST=$(HOST) PORT=$(PORT) ./intech-bot

image:
	@$(CTN) build -t $(REGISTRY)/intech-bot .

ctn-run: image
	@$(CTN) run \
		-e TKN=$(TKN) \
		-e HOST=$(HOST) \
		-e PORT=$(PORT) \
		-p $(PORT):$(PORT) \
		intech-bot

publish:
	@$(CTN) push $(REGISTRY)/intech-bot
