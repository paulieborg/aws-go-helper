#!/bin/bash

go run main.go -n MyTestStack -a provision -b essentials-tf-state
go run main.go -n MyTestStack -a provision -b essentials-tf-state

sed -i '' 's#10.0.0.0/16#10.1.0.0/16#' templates/test-params.json

go run main.go -n MyTestStack -a provision -b essentials-tf-state
go run main.go -n MyTestStack -a delete

git checkout -- templates/test-params.json
