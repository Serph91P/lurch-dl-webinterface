#!/usr/bin/env bash

OUTPUT_DIR="./dist"
VERSION=$(git describe --tags)

function gobuild {
    printf " * ${GOOS}\t${GOARCH}\t${OUTPUT_FILE} "
    go build -ldflags="-X 'main.Version=${VERSION}'" -o "${OUTPUT_DIR}/${OUTPUT_FILE}" && printf "\t✔\n"
}

echo "Building version ${VERSION} into ${OUTPUT_DIR}"

GOOS=windows GOARCH=386   OUTPUT_FILE=lurch-dl_${VERSION}_32bit.exe   gobuild
GOOS=windows GOARCH=amd64 OUTPUT_FILE=lurch-dl_${VERSION}_64bit.exe   gobuild
GOOS=linux   GOARCH=386   OUTPUT_FILE=lurch-dl_${VERSION}_linux_i386  gobuild
GOOS=linux   GOARCH=amd64 OUTPUT_FILE=lurch-dl_${VERSION}_linux_amd64 gobuild
GOOS=linux   GOARCH=arm   OUTPUT_FILE=lurch-dl_${VERSION}_linux_arm   gobuild
GOOS=linux   GOARCH=arm64 OUTPUT_FILE=lurch-dl_${VERSION}_linux_arm64 gobuild
