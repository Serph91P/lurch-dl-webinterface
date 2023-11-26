FROM docker.io/golang:1.21-alpine3.18 AS builder

ARG VERSION="development"

WORKDIR /usr/src/app
COPY . .
RUN go build -o /usr/local/bin/lurch-dl -ldflags="-X 'main.Version=${VERSION}'"

FROM docker.io/alpine:3.18

WORKDIR /output/
COPY --from=builder /usr/local/bin/lurch-dl /usr/local/bin/lurch-dl

RUN apk add --no-cache python3 py3-pip \
    && pip3 install flask gunicorn

COPY /web_ui/app.py /webapp/app.py
COPY /web_ui/templates /webapp/templates

COPY start.sh /start.sh
RUN chmod +x /start.sh

WORKDIR /webapp

ENTRYPOINT ["/start.sh"]