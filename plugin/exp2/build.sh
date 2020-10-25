#!/bin/bash
go build --buildmode=plugin -o plugin.so plugin.go
go run main.go
Hello, number 7
