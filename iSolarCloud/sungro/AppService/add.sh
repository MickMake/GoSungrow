#!/bin/bash

DEST="$1"
if [ "${DEST}" == "" ]
then
	echo "Need a destination endpoint"
	exit
fi

mkdir "${DEST}"
perl -pe "s/nullEndPoint/${DEST}/g" nullEndPoint/data.go > "${DEST}/data.go"
perl -pe "s/nullEndPoint/${DEST}/g" nullEndPoint/struct.go > "${DEST}/struct.go"

