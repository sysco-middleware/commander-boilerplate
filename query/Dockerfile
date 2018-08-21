# Builder
FROM syscomiddleware/commander-build AS builder

ARG GO_PROJECT_PATH=${GO_PROJECT_PATH}

WORKDIR /go/src/${GO_PROJECT_PATH}
COPY . .

RUN dep ensure
RUN go build -o build .

# Final
FROM alpine as final

ARG GO_PROJECT_PATH=${GO_PROJECT_PATH}

COPY --from=builder /go/src/${GO_PROJECT_PATH}/build ./
ENTRYPOINT ["./build"]
