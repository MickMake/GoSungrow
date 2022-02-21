#!/bin/bash

AREA="$1"
if [ "${AREA}" == "" ]
then
	echo "Need a destination area"
	exit
fi

NAME="$2"
if [ "${NAME}" == "" ]
then
	echo "Need a destination endpoint"
	exit
fi

URL="$3"
if [ "${URL}" == "" ]
then
	echo "Need an endpoint url"
	exit
fi



DIRNAME="$(realpath "${AREA}/${NAME}")"
TEMPLATE="$(realpath "../api/nullEndPoint/")"

if [ ! -d "${DIRNAME}" ]
then
	mkdir "${DIRNAME}"
else
	echo "Directory ${DIRNAME} exists"
	exit
fi

if [ ! -f "${DIRNAME}/data.go" ]
then
	perl -pe "s#nullEndPoint#${NAME}#g; s#%URL%#${URL}#g" ${TEMPLATE}/data.go > "${DIRNAME}/data.go"
else
	echo "File ${DIRNAME}/data.go exists"
fi

if [ ! -f "${DIRNAME}/struct.go" ]
then
	perl -pe "s/nullEndPoint/${NAME}/g" ${TEMPLATE}/struct.go > "${DIRNAME}/struct.go"
else
	echo "File ${DIRNAME}/struct.go exists"
fi

