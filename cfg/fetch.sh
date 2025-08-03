#!/bin/bash
host=$1
target=$2

# Check for required dependencies
for cmd in curl yq; do
    if ! command -v "$cmd" &> /dev/null; then
        echo "Error: $cmd is not installed"
        exit 1
    fi
done

curl -v "https://${host}/api/v2.0" \
  --compressed | yq -P > $target
