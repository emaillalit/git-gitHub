#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o tag-role-render-darwin-amd64

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tag-role-render-linux-amd64
