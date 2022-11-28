# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-alpine AS build
RUN apk update && apk --no-cache --update add build-base

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV GOOS=linux CGO_ENABLED=1
RUN go build -o /spaceapi-server

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /spaceapi-server /spaceapi-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/spaceapi-server"]
