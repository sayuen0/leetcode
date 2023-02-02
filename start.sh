#!/bin/bash
set -e
re='^[0-9]+$'
if ! [[ $1 =~ $re ]] ; then
   echo "error: Not a number" >&2; exit 1
fi
if [[ -z $2 ]] ; then
   echo "error: Empty Problem name" >&2; exit 1
fi

d=no_$(printf "%02d\n" $1)
f=$2
mkdir $d

echo "package main" > $d/$f.go
echo "package main" > $d/"${f}"_test.go
