#!/bin/bash

docker-compose run build

docker-compose run provision
docker-compose run provision

sed -i '' 's#10.0.0.0/16#10.1.0.0/16#' templates/test-params.json

docker-compose run provision
docker-compose run delete

git checkout -- templates/test-params.json
