#!/bin/bash

CMD=${1:-"web"}
OUTPUT=${2:-"main"}
OS=${3:-$(go env GOOS)}
ARCH=${4:-$(go env GOARCH)}

echo "Compiling $CMD for $OS/$ARCH into /bin/$OUTPUT..."

GOOS=$OS GOARCH=$ARCH \
    go build \
    -o bin/$OUTPUT \
    -ldflags="-X main.isBuild=true -X main.entry=$CMD" \
    main.go

if [ $? -eq 0 ]; then
    echo "Build successful"
else
    echo "Build failed"
    exit 1
fi