#!/bin/bash

openapi-generator-cli generate -i cfg/openapi.yaml -c cfg/config.yaml -o . -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
patch client.go < client.patch
go fmt ./...
