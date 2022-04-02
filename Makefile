BIN := "./bin/spa"
DOCKER_IMG="spa:develop"

#GIT_HASH := $(shell git log --format="%h" -n 1)
DATABASE_URL := "postgres://postgres:postgres@localhost:54321/spa?sslmode=disable"
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

buildx:
	#go build -v -o $(BIN) ./cmd/spa
	go build -v -o ./bin/spa ./cmd/spa

run: build
	$(BIN) -config ./configs/config.yaml

build-img:
	docker build \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .
#		--build-arg=LDFLAGS="$(LDFLAGS)" \


run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

migration:
	goose --dir=migrations postgres ${DATABASE_URL} up
