.PHONY: down up

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
	./script/run
