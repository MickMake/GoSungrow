#!/bin/bash

SUNGRO_GIT_TOKEN="$(jq -r '."git-token"' $HOME/.GoSungrow/config.json)"
SUNGRO_HOST="$(jq -r '.host' $HOME/.GoSungrow/config.json)"
SUNGRO_ID="$(jq -r '.id' $HOME/.GoSungrow/config.json)"
SUNGRO_PASSWORD="$(jq -r '.password' $HOME/.GoSungrow/config.json)"
SUNGRO_SECRET="$(jq -r '.secret' $HOME/.GoSungrow/config.json)"
SUNGRO_USER="$(jq -r '.user' $HOME/.GoSungrow/config.json)"
SUNGRO_GIT_REPO="$(jq -r '."git-repo"' $HOME/.GoSungrow/config.json)"
SUNGRO_GIT_DIR="/SUNGRO"
SUNGRO_DIFF_CMD="tkdiff"
# SUNGRO_GIT_TOKEN SUNGRO_HOST SUNGRO_ID SUNGRO_PASSWORD SUNGRO_SECRET SUNGRO_USER SUNGRO_DIFF_CMD

docker build -t gosungro \
	-f Dockerfile \
	--build-arg 'GO_REPO_TOKEN=glpat-tFrj4ZD7soVU2fqxuDMh' \
	--build-arg "SUNGRO_HOST=${SUNGRO_HOST}" \
	--build-arg "SUNGRO_USER=${SUNGRO_USER}" \
	--build-arg "SUNGRO_PASSWORD=${SUNGRO_PASSWORD}" \
	--build-arg "SUNGRO_ID=${SUNGRO_ID}" \
	--build-arg "SUNGRO_SECRET=${SUNGRO_SECRET}" \
	--build-arg "SUNGRO_GIT_REPO=${SUNGRO_GIT_REPO}" \
	--build-arg "SUNGRO_GIT_DIR=${SUNGRO_GIT_DIR}" \
	--build-arg "SUNGRO_GIT_TOKEN=${SUNGRO_GIT_TOKEN}" \
	--build-arg "SUNGRO_DIFF_CMD=${SUNGRO_DIFF_CMD}" \
	.

docker run --rm -it gosungro:latest /usr/local/bin/GoSungrow help
docker run --rm -it gosungro:latest /bin/sh

