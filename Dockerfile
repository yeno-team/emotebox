# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build

WORKDIR /api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY /controller ./controller
COPY /mock ./mock
COPY /pkg ./pkg
COPY /cmd ./cmd
COPY /docs ./docs
COPY /scripts ./scripts

RUN ["chmod", "+x", "./scripts/build.sh"]
RUN ./scripts/build.sh

CMD [ "/api/bin/emotebox" ]