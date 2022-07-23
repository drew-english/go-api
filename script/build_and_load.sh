#!/bin/bash

docker build -t go-api:local -f local_dockerfile .
kind load docker-image go-api:local
