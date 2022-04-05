#DATABASE_URL := "postgres://postgres:postgres@host.docker.internal:5432/spa?sslmode=disable"
DATABASE_URL := "postgresql://postgres:postgres@localhost:54321/spa?sslmode=disable"

buildx:
	go build -v -o ./bin/spa ./cmd/spa

run: build
	./bin/spa -config ./configs/config.yaml

build-test:
	docker-compose build

run-test:
	docker-compose up

build-prod:
	docker-compose -f docker-compose-prod.yaml build

run-prod:
	docker-compose -f docker-compose-prod.yaml up

test:
	go test -race ./internal/...

migration:
	goose --dir=migrations postgres ${DATABASE_URL} up
