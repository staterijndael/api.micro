#!/bin/bash
#title           :7_delete_images.sh
#description     :Remove all services images from local.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/))
readonly tag=1.0.0

for i in "${arr[@]}"
do
  docker rmi "deissh/api-micro-$i"
done
