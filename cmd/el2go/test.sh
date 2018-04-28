#!/usr/bin/env bash

for d in *
do
  if [ -d $d ]; then
    cd $d
    go build
  fi
done
