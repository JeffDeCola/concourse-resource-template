#!/bin/sh
# concourse-resource-template deploy.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "************************************************************************"
    echo "* deploy.sh -debug (START) *********************************************"
    echo "************************************************************************"
    # set -x enables a mode of the shell where all executed commands
    # are printed to the terminal.
    set -x
    echo " "
else
    echo "************************************************************************"
    echo "* deploy.sh (START) ****************************************************"
    echo "************************************************************************"
    echo " "
fi

echo "Deploy to Docker"
docker run --name concourse-resource-template -dit jeffdecola/concourse-resource-template
echo " "

echo "************************************************************************"
echo "* deploy.sh (END) ******************************************************"
echo "************************************************************************"
echo " "
