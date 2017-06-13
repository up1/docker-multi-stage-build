#!/bin/sh
echo "Build step 01"

docker image build -t my_image:build . -f Dockerfile_step1

docker create --name extract my_image:build
docker cp extract:/go/src/app ./app
docker rm -f extract

echo "Build step 02"
docker image build --no-cache -t my_image:latest . -f Dockerfile_step2
