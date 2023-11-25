#!/usr/bin/env bash

VERSION=$(git describe --tags)

echo "Building version ${VERSION} as container..."

podman build -t "lurch-dl:${VERSION}" --build-arg="VERSION=${VERSION}" .
