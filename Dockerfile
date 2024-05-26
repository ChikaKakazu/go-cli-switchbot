FROM golang:1.22.3-alpine3.19

RUN apk update && apk add git curl unzip

ENV GOPATH /go

RUN mkdir /go/app
COPY . /go/app

WORKDIR /go/app
