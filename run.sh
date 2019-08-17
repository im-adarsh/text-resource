#!/bin/bash
docker build . -t text-resource
docker run -i -t -p 8080:8080 text-resource go run text-resource/text-resource/cmd/server/main.go