#!/bin/bash
#title           :0create_docs.sh
#description     :Build Swagger and other docs.
#author		     :deissh
#version         :0.1
#=====================================

readonly -a arr=($(ls -d service-*/ | xargs -n 1 basename))

for i in "${arr[@]}"
do
  $GOPATH/bin/swag init --output ${i}/docs
done