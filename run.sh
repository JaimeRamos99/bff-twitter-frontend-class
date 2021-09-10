#!/bin/sh
docker build -t bff-go .
docker run --rm -d -p 8080:8080 bff-go