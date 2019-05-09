#!/bin/bash
#title           :2_push_images.sh
#description     :Push docker images.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d */))
readonly tag=1.4.0

for i in "${arr[@]}"
do
  docker push "deissh/api-micro-$i:$tag"
done
