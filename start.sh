#!/bin/bash
set -e

if [[ -z $1 ]] ; then
   echo "error: Empty Problem name" >&2; exit 1
fi

# Combine all arguments into a single string
input="${@}"
first=$(echo "$input" | cut -d'.' -f1)
second=$(echo "$input" | cut -d'.' -f2- | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//' | tr '[:upper:]' '[:lower:]' | sed -r 's/([A-Z])/_\L\1/g' | sed 's/^_//' | tr ' ' '_')

d=no_$(printf "%03d\n" $first)
f=$second

# Create problem directory
mkdir -p $d

echo "package main" > $d/$f.go
echo "package main" > $d/"${f}"_test.go

# Create example directory
mkdir -p $d/example

echo "package main" > $d/example/$f.go
echo "package main" > $d/example/"${f}"_test.go
