#!/bin/bash

kind create cluster --config kind.yml
script/build_and_load.sh
helm install go-api helm/go-api
