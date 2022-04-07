FROM golang:1.16-alpine
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
#RUN go mod download

COPY . .

RUN go install -v ./...


RUN GO_ENABLED=1 go build -o ./bin/spa ./cmd/spa

EXPOSE 8081

CMD ./bin/spa -config ./configs/config.yaml