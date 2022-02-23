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

if [ "${NAME}" == "login" ]
then
	echo "Refusing to modify login EndPoint"
	exit
fi

CMD="$(realpath $0)"
CMD="$(dirname ${CMD})"

DIRNAME="$(realpath "${AREA}/${NAME}/")"
DIRNAME="$(echo "${DIRNAME}" | perl -pe "s#^${CMD}/##")"
TEMPLATE="$(realpath "api/nullEndPoint/")"
TEMPLATE="$(echo "${TEMPLATE}" | perl -pe "s#^${CMD}/##")"

# echo "Template: ${TEMPLATE}"
echo "EndPoint: ${DIRNAME}"

BC="$(cksum "${DIRNAME}/struct.go")"

if [ ! -d "${DIRNAME}" ]
then
	echo "Directory ${DIRNAME} does not exist."
	exit
fi

if [ ! -f "${DIRNAME}/struct.go" ]
then
	echo "File ${DIRNAME}/struct.go does not exist."
else
	perl -pe "s/nullEndPoint/${NAME}/g" ${TEMPLATE}/struct.go > "${DIRNAME}/struct.go"
	AC="$(cksum "${DIRNAME}/struct.go")"

	echo "${BC}"
	echo "${AC}"
fi

