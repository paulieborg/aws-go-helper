#!/bin/bash

set -e -o pipefail

echo '{
  "VpcID": "vpc-afe650cb",
  "InPort": "80",
  "OutPort": "80"
}' > templates/test-params.json

go run main.go -n MyTestStack -a provision #-b myob-dont-panic-test
go run main.go -n MyTestStack -a provision #-b myob-dont-panic-test

if [[ $(uname -s) == "Darwin" ]]; then
    sed -i '' 's/80/443/' templates/test-params.json
else
    sed -i 's/80/443/' templates/test-params.json
fi

go run main.go -n MyTestStack -a provision #-b myob-dont-panic-test
go run main.go -n MyTestStack -a delete

rm templates/test-params.json
