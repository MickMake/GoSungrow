#!/bin/bash

DATE="$(date +'%Y%m%d')"

./get.sh queryUserCurveTemplateData
mv queryUserCurveTemplateData.response queryUserCurveTemplateData-${DATE}.response

echo '"Key","Date","Value"' > queryUserCurveTemplateData-${DATE}.csv

cat queryUserCurveTemplateData-${DATE}.response | \
jq -r '
.result_data.points_data."1129147_11_0_0" |
 with_entries(if (.key | test("p[0-9]+$")) then ( {key: .key, value: .value } ) else empty end ) |
 .[] | 
"\"" + .point_name + "\"," + (.data_list[] | ((.time_stamp | strptime("%Y%m%d%H%M%S") | strftime("%Y/%m/%d %H:%M:%S")) + "," + .value))' >> queryUserCurveTemplateData-${DATE}.csv

