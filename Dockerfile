FROM golang:1.18-alpine AS builder
RUN mkdir -p /go/src/gitlab.com/go-rabbitmq-consumer-app

ENV GOPATH /go
WORKDIR /go/src/gitlab.com/go-rabbitmq-consumer-app

ADD go.mod /go/src/gitlab.com/go-rabbitmq-consumer-app
ADD go.sum /go/src/gitlab.com/go-rabbitmq-consumer-app
ADD . /go/src/gitlab.com/go-rabbitmq-consumer-app

RUN go build

FROM alpine

RUN mkdir -p /app
COPY --from=builder /go/src/gitlab.com/go-rabbitmq-consumer-app/configs/config.dev.yaml /app/
COPY --from=builder /go/src/gitlab.com/go-rabbitmq-consumer-app/go-rabbitmq-consumer-app /app/
RUN chmod +x /app/go-rabbitmq-consumer-app
WORKDIR /app
ENTRYPOINT ["/app/go-rabbitmq-consumer-app"]