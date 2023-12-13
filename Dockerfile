# Stage 1: Build lurch-dl from the latest GitHub repository
FROM golang:1.21-alpine3.18 AS builder

# Installiere Git
RUN apk add --no-cache git

# Klone das Repository
RUN git clone https://github.com/ChaoticByte/lurch-dl.git /usr/src/lurch-dl

# Wechsle in das Repository-Verzeichnis
WORKDIR /usr/src/lurch-dl

# Baue lurch-dl
RUN go build -o /usr/local/bin/lurch-dl

# Stage 2: Final Image
FROM alpine:3.18

# Installiere Python und Quart
RUN apk add --no-cache python3 py3-pip \
    && pip3 install quart uvicorn

# Setze das Arbeitsverzeichnis
WORKDIR /webapp

# Kopiere lurch-dl vom Builder
COPY --from=builder /usr/local/bin/lurch-dl /usr/local/bin/lurch-dl

# Startskript für Quart-Server (und möglicherweise weitere Initialisierung)
COPY start.sh /start.sh
RUN chmod +x /start.sh

# ENTRYPOINT auf das Startskript setzen
ENTRYPOINT ["/start.sh"]
