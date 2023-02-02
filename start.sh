#!/bin/bash
set -e
re='^[0-9]+$'
if ! [[ $1 =~ $re ]] ; then
   echo "error: Not a number" >&2; exit 1
fi

d=no_$(printf "%02d\n" $1)
mkdir $d

echo "package main" > $d/main.go
echo "package main" > $d/main_test.go
