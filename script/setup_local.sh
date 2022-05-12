#!/bin/bash

kind create cluster --config kind.yml
helm install go-api helm/go-api
