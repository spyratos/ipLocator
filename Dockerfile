ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY . .
RUN go mod download


RUN go build -o ./app ./httpd/main.go
FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY --from=builder /api/app .

EXPOSE 8080

ENTRYPOINT ["./app"]