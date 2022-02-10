#!/bin/bash

DATE="$1"
if [ "${DATE}" == "" ]
then
	DATE="$(date +'%Y%m%d')"
fi

./get.sh queryMutiPointDataList ${DATE}

echo '"Key","Date","Value"' > queryMutiPointDataList-${DATE}.csv

cat queryMutiPointDataList-${DATE}.response | \
jq -r '
.result_data."1129147_11_0_0" | 
 with_entries(if (.key | test("p[0-9]+$")) then ( {key: .key, value: .value } ) else empty end ) | 
 to_entries | 
 .[] | 
 "\"" + .key + "\"," + (.value | to_entries | .[] | "\"" + (.key | strptime("%Y%m%d%H%M%S") | strftime("%Y/%m/%d %H:%M:%S")) + "\"," + .value)' >> queryMutiPointDataList-${DATE}.csv

