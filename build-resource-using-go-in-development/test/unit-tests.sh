#!/bin/sh -e
# concourse-resource-template unit-tests.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "************************************************************************"
    echo "* unit-tests.sh -debug (START) *****************************************"
    echo "************************************************************************"
    # set -x enables a mode of the shell where all executed commands
    # are printed to the terminal.
    set -x
    echo " "
else
    echo "************************************************************************"
    echo "* unit-tests.sh (START) ************************************************"
    echo "************************************************************************"
    echo " "
fi

echo "If you need to generate _test files run"
echo "   gotests -w -all main.go"
echo " "

echo "cd up to code level"
cd ..
echo " "

echo "go test -cover"
echo "   -cover shows the percentage coverage"
echo "   Put results in /test/test_coverage.txt file"
echo " "
go test -cover ./... | tee test/test_coverage.txt
echo " "

echo "************************************************************************"
echo "* unit-tests.sh (END) **************************************************"
echo "************************************************************************"
echo " "
