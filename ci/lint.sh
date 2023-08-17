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

echo "Linting frontend..."
pushd frontend
golangci-lint run ./...
popd

echo "Linting backend..."
pushd backend
golangci-lint run ./...
popd