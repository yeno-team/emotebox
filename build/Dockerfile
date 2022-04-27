# syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build

WORKDIR /api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN bash ./scripts/build.sh


FROM build as dev

# live reload with air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
EXPOSE 8080
CMD [ "air" ]




CMD [ "/emotebox-server" ]