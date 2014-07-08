.PHONY: down up worker

test:
	godep go test ./...

integration:
	godep go test -tags=integration ./...

down:
	goose -env development down
	goose -env test down

up:
	goose -env development up
	goose -env test up

build:
	godep go build -o build/server github.com/remind101/shipr/server

run: build
	build/server
