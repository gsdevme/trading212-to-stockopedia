FROM alpine:3

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /tmp

ENTRYPOINT ["/usr/local/bin/trading212-to-stockopedia"]
COPY trading212-to-stockopedia /usr/local/bin