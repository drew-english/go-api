#!/bin/bash

docker build -t go-api:local -f Dockerfile.local .
kind load docker-image go-api:local
