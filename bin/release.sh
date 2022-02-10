#!/bin/bash

VERSION="$1"
if [ -z "${VERSION}" ]
then
	echo "Need a version number. EG: v1.42"
	exit 1
fi

# vi defaults/const.go 
git tag -d "${VERSION}"
rm -rf dist/

git add .
git commit -m "Committed ${VERSION}"

git tag "${VERSION}"
goreleaser build

git add .
git commit -m "Released ${VERSION}"

git push
