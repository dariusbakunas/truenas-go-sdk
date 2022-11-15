host=$1
target=$2

curl "https://${host}/api/v2.0" \
  --compressed | yq -P > $target