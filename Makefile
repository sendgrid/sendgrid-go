.PHONY: test install

install:
	go get -t -v ./...

test: install
	./go.coverage.sh
	go test -cover -v -race ./... | grep -v -E '/vendor|/examples|/docker'
