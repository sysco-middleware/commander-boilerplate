# Builder
FROM golang:1.10-alpine AS builder

ENV GO_PROJECT_PATH=query

WORKDIR /go/src/${GO_PROJECT_PATH}
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o build .

# Final
FROM alpine

ENV GO_PROJECT_PATH=query

COPY --from=builder /go/src/${GO_PROJECT_PATH}/build ./
ENTRYPOINT ["./build"]
