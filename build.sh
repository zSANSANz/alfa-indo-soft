#!/bin/bash

branch=$(git symbolic-ref --short HEAD)

case $branch in
    "master" | *)
    prefix="master"
    ;;
    "staging")
    prefix="staging"
    ;;
    "development")
    prefix="dev"
    ;;
esac

docker build . -t backend-$prefix

docker stop  backend-$prefix

docker container prune -f

docker run --restart always -d -p 8585:8585 --name  backend-$prefix  backend-$prefix

docker system prune -af