#!/bin/bash
go build -o wowtoken-api ./cmd/api
docker build -t wowtoken-api .
docker run -d -p 8080:8080 wowtoken-api