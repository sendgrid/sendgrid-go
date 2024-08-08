.PHONY: test install test-integ test-docker

install:
	go get -t -v ./...

test:
	./go.coverage.sh
	bash -c 'diff -u <(echo -n) <(gofmt -d -s .)'

test-integ: test

version ?= latest
test-docker:
	docker compose up --build --parallel --force-recreate --abort-on-container-exit --remove-orphans
