#!/bin/bash
# resource-template unit-test.sh

set -e -x

# The code is located in /resource-template
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./resource-template to $GOPATH/src/github.com/JeffDeCola/resource-template
mkdir -p src/github.com/JeffDeCola/
cp -R ./resource-template src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/JeffDeCola/resource-template

# RUN unit_tests
# go test -v -cover ./...