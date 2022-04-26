# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN go build -o /emotebox-server
EXPOSE 8080

CMD [ "/emotebox-server" ]