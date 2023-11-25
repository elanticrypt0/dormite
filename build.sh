#!/bin/bash

# binaries
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s"
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s"

rm -rf ./build
mkdir ./build

mv dormite ./build
mv dormite.exe ./build