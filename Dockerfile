# syntax=docker/dockerfile:1

FROM golang:1.16-buster as build

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

## DEVELOPMENT STAGE

FROM build as dev

# install air on dev stage for development
RUN go get github.com/cosmtrek/air


## PRODUCTION STAGE

FROM gcr.io/distroless/base-debian10 as prod

WORKDIR /
COPY --from=build /api/bin/emotebox /api/bin/emotebox

EXPOSE 8080
USER nonroot:nonroot

# ENTRYPOINT [ "/api/bin/emotebox" ]