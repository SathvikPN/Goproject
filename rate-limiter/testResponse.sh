#!/bin/bash

URL="http://localhost:8080/"
for i in {1..10}
do
  for j in {1..4}
  do
    echo -n "$i.$j: " 
    curl -s $URL 
    if [[ $? -ne 0 ]]; then 
      echo ""
    fi
  done
  sleep 1s
done
