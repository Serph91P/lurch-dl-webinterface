# Stage 1: Build lurch-dl
FROM docker.io/golang:1.21-alpine3.18 AS builder
ARG VERSION="development"
WORKDIR /usr/src/app
COPY . .
RUN go build -o /usr/local/bin/lurch-dl -ldflags="-X 'main.Version=${VERSION}'"

# Stage 2: Final Image
FROM docker.io/alpine:3.18
WORKDIR /output/

# Kopiere lurch-dl vom Builder
COPY --from=builder /usr/local/bin/lurch-dl /usr/local/bin/lurch-dl

# Installiere Python, Flask, Gunicorn, Uvicorn und Quart
RUN apk add --no-cache python3 py3-pip \
    && pip3 install flask gunicorn uvicorn quart


# Kopiere die Flask-App und Templates in das Image
COPY /web_ui/app.py /webapp/app.py
COPY /web_ui/templates /webapp/templates

# Setze das Arbeitsverzeichnis
WORKDIR /webapp

# Startskript für Flask-Server (und möglicherweise weitere Initialisierung)
COPY start.sh /start.sh
RUN chmod +x /start.sh

# ENTRYPOINT auf das Startskript setzen
ENTRYPOINT ["/start.sh"]
