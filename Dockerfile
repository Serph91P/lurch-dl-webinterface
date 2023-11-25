FROM docker.io/golang:1.21-alpine3.18 AS builder

ARG VERSION="development"

WORKDIR /usr/src/app
COPY . .
RUN go build -o /usr/local/bin/lurch-dl -ldflags="-X 'main.Version=${VERSION}'"

FROM docker.io/alpine:3.18

WORKDIR /output/
COPY --from=builder /usr/local/bin/lurch-dl /usr/local/bin/lurch-dl

ENTRYPOINT ["/usr/local/bin/lurch-dl"]
