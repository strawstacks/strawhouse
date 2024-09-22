# Strawhouse
A file store engine focused on lightweight indexing with access control.

## Installation
1. Install dependencies
    ```bash
    brew install protobuf
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
   
2. Source commands
    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```