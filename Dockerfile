FROM alpine:3

WORKDIR /tmp

ENTRYPOINT ["/usr/local/bin/trading212-to-stockopedia"]
COPY trading212-to-stockopedia /usr/local/bin