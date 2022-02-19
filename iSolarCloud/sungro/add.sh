#!/bin/bash

DEST="$1"
if [ "${DEST}" == "" ]
then
	echo "Need a destination endpoint"
	exit
fi

mkdir "${DEST}"

if [ ! -f "${DEST}/data.go" ]
then
	perl -pe "s/nullEndPoint/${DEST}/g" nullEndPoint/data.go > "${DEST}/data.go"
else
	echo "${DEST}/data.go exists"
fi

if [ ! -f "${DEST}/struct.go" ]
then
	perl -pe "s/nullEndPoint/${DEST}/g" nullEndPoint/struct.go > "${DEST}/struct.go"
else
	echo "${DEST}/struct.go exists"
fi

