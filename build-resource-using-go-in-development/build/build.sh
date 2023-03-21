#!/bin/sh -e
# concourse-resource-template build.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "************************************************************************"
    echo "* build.sh -debug (START) **********************************************"
    echo "************************************************************************"
    # set -x enables a mode of the shell where all executed commands
    # are printed to the terminal.
    set -x
    echo " "
else
    echo "************************************************************************"
    echo "* build.sh (START) *****************************************************"
    echo "************************************************************************"
    echo " "
fi

echo "cd to where go code is"
echo "cd .."
cd ..
echo " " 

echo "Build your docker image using Dockerfile"
echo "NOTE: The binary is built using this step"
echo "docker build -f build/Dockerfile -t jeffdecola/concourse-resource-template ."
docker build -f build/Dockerfile -t jeffdecola/concourse-resource-template .
echo " "

echo "Check Docker Image size"
echo "docker images jeffdecola/concourse-resource-template:latest"
docker images jeffdecola/concourse-resource-template:latest
echo " "

echo "Useful commands:"
echo "     docker run --name concourse-resource-template -dit jeffdecola/concourse-resource-template"
echo "     docker exec -i -t concourse-resource-template /bin/bash"
echo "     docker logs concourse-resource-template"
echo " "

echo "************************************************************************"
echo "* build.sh (END) *******************************************************"
echo "************************************************************************"
echo " "
