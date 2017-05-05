[![Build status](https://badge.buildkite.com/0b1e1cd70a91f5e6b694877f4b86d7e7d74b5e2ec6ce427b28.svg)](https://buildkite.com/myob/aws-go-helpers)

# aws-go-helper
##Helpers using AWS GoSDK

    docker-compose run --rm build
    docker-compose run --rm run

##Prerequisites

Using https://github.com/Masterminds/glide

On Mac: `brew install glide`

Setup project like so:
```
mkdir -p ${HOME}/go/src/github.com/paulieborg
export GOPATH=${HOME}/go
cd ${HOME}/go/src/github.com/paulieborg
git clone git@github.com:paulieborg/aws-go-helper.git
cd aws-go-helper
glide install
```

