#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=1.0.0

for i in "${arr[@]}"
do
  docker build -t "deissh/api-micro-$i:$tag" --build-arg service_name=${i} .
done

docker image ls | grep 'deissh/api-micro-'