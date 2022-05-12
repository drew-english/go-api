#!/bin/bash

docker build -t go-api:local .
kind load docker-image go-api:local
