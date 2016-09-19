#!/bin/bash

set -e -x
 
export GOPATH=$PWD
mkdir -p src/github.com/JeffDeCola/
cp -R ./resource-template src/github.com/JeffDeCola/.

go test -v -cover github.com/JeffDeCola/resource-template/...