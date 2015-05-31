#!/bin/sh
for dir in */; do
  cd $dir
  go install
  cd -
done
