#!/bin/bash

VERSION="$(awk '/BinaryVersion/{gsub("\"",""); print"v"$3}' defaults/const.go)"
if [ -z "${VERSION}" ]
then
	echo "Unknown version."
	exit 1
fi

echo -n "Checking git tag \"${VERSION}\"  - "
TAG="$(git describe --tags "${VERSION}")"
if [ ! -z "${TAG}" ]
then
	echo "Version ${VERSION} already exists."
	echo "Delete with:"
	echo "git tag -d ${VERSION}"
  echo ""
	echo "OR modify defaults/const.go file."
  echo ""
	echo "OR hit enter to just commit and push."
  echo ""
  echo -n "OK? (Ctrl-C to terminate) "
  read OK

  git add .
  git commit -m "Committed ${VERSION}"
  git push

  exit 1
fi

echo ""
echo "Releasing version ${VERSION}"
echo ""
echo -n "OK? (Ctrl-C to terminate) "
read OK

git add .
git commit -m "Committed ${VERSION}"
git tag "${VERSION}"
git push

goreleaser release --rm-dist

git add .
git commit -m "Released ${VERSION}"
git push

