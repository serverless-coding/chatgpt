set -euxo pipefail

go build -o ./functions/chatgpt4 ./functions/main.go
ls -al
GOBIN=$(pwd)/functions go install ./...
chmod +x "$(pwd)"/functions/*
go env