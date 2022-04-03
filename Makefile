BIN := "./bin/spa"
DOCKER_IMG="spa:develop"

#DATABASE_URL := "postgres://postgres:postgres@host.docker.internal:5432/spa?sslmode=disable"
DATABASE_URL := "postgresql://postgres:postgres@localhost:54321/spa?sslmode=disable"

buildx:
	#go build -v -o $(BIN) ./cmd/spa
	go build -v -o ./bin/spa ./cmd/spa

run: build
	$(BIN) -config ./configs/config.yaml

build-img:
	docker build \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

migration:
	goose --dir=migrations postgres ${DATABASE_URL} up
