#!/usr/bin/env bash

VERSION=$(git describe --tags)

echo "Building version ${VERSION}"

# linux i386
printf "... for linux i386 "
GOOS=linux GOARCH=386 go build -ldflags="-X 'main.Version=${VERSION}'" -o dist/lurch-dl_${VERSION}_linux_i386 && printf "\t✔\n"

# linux amd64
printf "... for linux amd64 "
GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=${VERSION}'" -o dist/lurch-dl_${VERSION}_linux_amd64 && printf "\t✔\n"

# linux arm
printf "... for linux arm "
GOOS=linux GOARCH=arm go build -ldflags="-X 'main.Version=${VERSION}'" -o dist/lurch-dl_${VERSION}_linux_arm && printf "\t✔\n"

# linux arm64
printf "... for linux arm64 "
GOOS=linux GOARCH=arm64 go build -ldflags="-X 'main.Version=${VERSION}'" -o dist/lurch-dl_${VERSION}_linux_arm64 && printf "\t✔\n"
