.PHONY: down up worker

test:
	godep go test ./...

down:
	goose -env development down
	goose -env test down

up:
	goose -env development up
	goose -env test up

install:
	godep go install

run: install
	shipr

worker:
	cd worker && make install
