#!/bin/bash

NAME="$1"

perl -pe 'BEGIN{use Env qw(NAME);}s/PACKAGE/$NAME/g;$n=ucfirst($NAME);s/STRUCTURE/$n/g' template.go > ../${NAME}/${NAME}.go 
perl -pe 'BEGIN{use Env qw(NAME);}s/PACKAGE/$NAME/g;$n=ucfirst($NAME);s/STRUCTURE/$n/g' template_struct.go > ../${NAME}/${NAME}_struct.go 
perl -pe 'BEGIN{use Env qw(NAME);}s/PACKAGE/$NAME/g;$n=ucfirst($NAME);s/STRUCTURE/$n/g' templateResponse.go > ../${NAME}/${NAME}Response.go 

perl -pe 'BEGIN{use Env qw(NAME);}s/PACKAGE/$NAME/g;$n=ucfirst($NAME);s/STRUCTURE/$n/g' ../TEMPLATE_sungro.go.txt > ../${NAME}_sungro.go 

