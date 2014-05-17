test:
	${GOPATH}/bin/godep go test ./...

install:
	${GOPATH}/bin/godep go install

run: install
	./script/run
