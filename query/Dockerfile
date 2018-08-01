# Builder
FROM golang:1.10-alpine AS builder

ARG GO_PROJECT_PATH=${GO_PROJECT_PATH}

WORKDIR /go/src/${GO_PROJECT_PATH}
COPY . .

RUN apk update && apk upgrade && \
    apk add git openssh

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v
RUN go build -o build .

# Final
FROM alpine as final

ARG GO_PROJECT_PATH=${GO_PROJECT_PATH}

COPY --from=builder /go/src/${GO_PROJECT_PATH}/build ./
ENTRYPOINT ["./build"]
