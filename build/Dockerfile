# Собираем в гошке
FROM golang:1.16.2 as build

ENV BIN_FILE /opt/spa/spa-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
#ARG LDFLAGS
RUN CGO_ENABLED=1 go build -o ${BIN_FILE} cmd/spa/*
#        -ldflags "$LDFLAGS" \

# На выходе тонкий образ
FROM alpine:3.15

LABEL ORGANIZATION="test"
LABEL SERVICE="test"
LABEL MAINTAINERS="test@test.test"

ENV BIN_FILE "/opt/spa/spa-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/spa-test/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}