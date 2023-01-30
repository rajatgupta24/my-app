FROM golang:1.19.4-alpine3.17 as builder

WORKDIR /app

COPY . /app

RUN go build .

ENTRYPOINT [" /app/go-app "]
