# syntax = docker/dockerfile:1

FROM golang:1.24.4-alpine AS build
RUN apk add build-base
ARG CGO_ENABLED=1
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o bot ./cmd/bot

FROM alpine:latest
COPY --from=build /app/bot /bot
RUN mkdir -p data
ENTRYPOINT ["/bot"]