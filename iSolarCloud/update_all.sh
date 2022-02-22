#!/bin/bash

gfind sungro/ -mindepth 3 -name struct.go -printf '%h\n' | egrep -v '/login/' | perl -pe 's#^sungro/(\w+)/(\w+)#./update.sh $1 $2#' > /tmp/z

bash /tmp/z

