#!/bin/bash

go mod tidy
reflex -r '(\.go$|go\.mod)' -s go run cmd/faves4/main.go
