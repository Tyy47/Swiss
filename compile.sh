#!/usr/bin/env bash

# Build for Windows
GOOS=windows go build -o release/swiss.exe

# Build for Linux
env GOOS=linux GOARCH=amd64 go build -o release/swiss
