#!/bin/bash
set -e
if [[ -z $1 ]] ; then
   echo "error: Empty Problem name" >&2; exit 1
fi

input="$1"
first=$(echo "$input" | cut -d'.' -f1)
second=$(echo "$input" | cut -d'.' -f2 | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//' | tr '[:upper:]' '[:lower:]' | sed -r 's/([A-Z])/_\L\1/g' | sed 's/^_//' | tr ' ' '_')

d=no_$(printf "%03d\n" $first)
f=$second

if [[ ! -d $d ]] ; then
   echo "error: Directory $d does not exist" >&2; exit 1
fi

mkdir -p $d/example

echo "package main" > $d/example/$f.go
echo "package main" > $d/example/"${f}"_test.go
