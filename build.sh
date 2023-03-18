#!/bin/bash
set -euxo pipefail

mkdir -p dist
copy ./public/index.html ./dist/index.html

go build -o ./functions/chatgpt4 ./functions/main.go
ls -al
GOBIN=$(pwd)/functions go install ./...

ls -al ./functions/
chmod +x "$(pwd)"/functions/*
go env

ls -al ./dist