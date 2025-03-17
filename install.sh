#!/bin/bash

go install google.golang.org/grpc@v1.32.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install github.com/cloudflare/cfssl/cmd/cfssl@v1.6.5
go install github.com/cloudflare/cfssl/cmd/cfssljson@v1.6.5
