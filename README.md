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

