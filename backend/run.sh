#!/bin/bash
export PATH=$PATH:$(go env GOPATH)/bin
swag init -g ./server/main/main.go
go run ./server/main/main.go