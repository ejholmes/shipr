.PHONY: down up build test frontend

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
	godep go build -o build/shipr github.com/remind101/shipr/server/shipr

frontend:
	cd frontend && npm install && gulp

run: build
	build/shipr
