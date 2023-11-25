#!/usr/bin/env bash

# Usage:
## Docker
# sudo ./build-container.sh
## Podman
# ./build-container.sh podman

VERSION=$(git describe --tags)

echo "Building version ${VERSION} as container..."

if [[ "$1" == "podman" ]]; then
    podman build -t "lurch-dl:${VERSION}" --build-arg="VERSION=${VERSION}" .
else
    docker build -t "lurch-dl:${VERSION}" --build-arg="VERSION=${VERSION}" .
fi
