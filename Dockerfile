# syntax = docker/dockerfile:1

FROM golang:1.24.4-alpine AS build
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build ./cmd/bot

FROM scratch
COPY --from=build /app/bot /bot
ENTRYPOINT ["/bot"]