FROM golang:1.10-alpine

ARG LIBRESSL_VERSION=${LIBRESSL_VERSION}
ARG LIBRDKAFKA_VERSION=${LIBRDKAFKA_VERSION}

# Install librdkafka
RUN apk update && apk upgrade && \
    apk add libressl${LIBRESSL_VERSION}-libcrypto libressl${LIBRESSL_VERSION}-libssl --update-cache --repository http://nl.alpinelinux.org/alpine/edge/main && \
    apk add librdkafka=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add librdkafka-dev=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

RUN go get -u github.com/golang/dep/cmd/dep
