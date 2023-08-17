#!/bin/bash
set -e
if ! command -v go &> /dev/null; then
    echo "Go could not be found, installing..."
    wget https://dl.google.com/go/go1.17.linux-amd64.tar.gz
    tar -xvf go1.17.linux-amd64.tar.gz
    sudo mv go /usr/local
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
fi

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

which golangci-lint || echo "golangci-lint is not installed or not in PATH"

echo "Linting frontend..."
ls -la
$(go env GOPATH)/bin/golangci-lint run /frontend/...
echo "Linting backend..."
$(go env GOPATH)/bin/golangci-lint run /backend/...
