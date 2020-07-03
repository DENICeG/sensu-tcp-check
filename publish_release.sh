#!/bin/bash

if [[ -z "$1" ]]; then 
  echo "need tag/version in format v1.x.y"
  exit 1
else
  TAG=$1
fi

CGO_ENABLED=0 go build -o bin/sensu-tcp-check cmd/sensu-tcp-check/main.go
tar czf sensu-tcp-check_${TAG}_linux_amd64.tar.gz bin/

sha512sum sensu-tcp-check_${TAG}_linux_amd64.tar.gz > sensu-tcp-check_${TAG}_sha512_checksums.txt
SHA_HASH_ONLY=$(cut -d " " -f 1 sensu-tcp-check_${TAG}_sha512_checksums.txt)

sed "s/__TAG__/${TAG}/g" sensu/asset_template.tpl > sensu/asset.yaml
sed -i "s/__SHA__/${SHA_HASH_ONLY}/g" sensu/asset.yaml

mkdir -p artifacts
rm -f artifacts/*
mv sensu-tcp-check_${TAG}_linux_amd64.tar.gz sensu-tcp-check_${TAG}_sha512_checksums.txt artifacts/

git add .
git commit
git tag $TAG
git push && git push --tags
