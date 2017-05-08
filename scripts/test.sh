#!/bin/bash

set -e -o pipefail

tempfile=$(mktemp)

echo '{
  "VpcID": "vpc-afe650cb",
  "InPort": "80",
  "OutPort": "80"
}' > ${tempfile}

go run main.go -n MyTestStack -a provision -p ${tempfile} #-b myob-dont-panic-test
go run main.go -n MyTestStack -a provision -p ${tempfile} #-b myob-dont-panic-test

if [[ $(uname -s) == "Darwin" ]]; then
    sed -i '' 's/80/443/' ${tempfile}
else
    sed -i 's/80/443/' ${tempfile}
fi

go run main.go -n MyTestStack -a provision -p ${tempfile} #-b myob-dont-panic-test
go run main.go -n MyTestStack -a delete
