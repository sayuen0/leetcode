#!/bin/bash
set -e
if [[ -z $1 ]] ; then
   echo "error: Empty Problem name" >&2; exit 1
fi


input="$1"
first=$(echo "$input" | cut -d'.' -f1)
second=$(echo "$input" | cut -d'.' -f2 | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//' |tr '[:upper:]' '[:lower:]' | sed 's/\b[[:upper:]]/\L&/g' | tr ' ' '_')

d=no_$(printf "%02d\n" $first)
f=$second

mkdir $d

echo "package main" > $d/$f.go
echo "package main" > $d/"${f}"_test.go
