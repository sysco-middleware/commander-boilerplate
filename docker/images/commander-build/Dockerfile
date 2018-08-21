FROM golang:1.10-alpine

RUN apk update && apk upgrade && \
    apk add git

RUN go get -u github.com/golang/dep/cmd/dep
