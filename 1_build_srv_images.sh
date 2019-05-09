#!/bin/bash
#title           :1_build_srv_images.sh
#description     :Build docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))
readonly tag=latest

for i in "${arr[@]}"
do
  docker build -t "deissh/api-micro-$i:$tag" -f ./${i}/Dockerfile .
  echo "$i done."
done

docker image ls | grep 'deissh/api-micro-'