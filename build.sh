#!/bin/bash
rm -rf build
mkdir build
go clean
env GOOS=linux go build -o build/tahara_client.linux
env GOOS=windows go build -o build/tahara_client.exe
env GOOS=darwin go build -o build/tahara_client.darwin
