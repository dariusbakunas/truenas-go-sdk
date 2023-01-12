#!/bin/bash

rm -rf angelfish
openapi-generator-cli generate -i cfg/angelfish.yaml -c cfg/angelfish_config.yaml -o ./angelfish -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
rm -rf angelfish/test # seems to be broken
patch angelfish/client.go < client.patch
mv angelfish/go.mod .
mv angelfish/go.sum .
cd angelfish || exit
go mod tidy
go fmt ./...
cd ..