#!/bin/bash

gfind AliSmsService AppService MttvScreenService PowerPointService WebAppService WebIscmAppService -mindepth 2 -name struct.go -printf '%h\n' | egrep -v '/login/' | perl -pe 's#^(\w+)/(\w+)#./update.sh $1 $2#' > /tmp/z

bash /tmp/z

