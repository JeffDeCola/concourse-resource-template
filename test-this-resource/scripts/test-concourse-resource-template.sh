#!/bin/sh
# concourse-resource-template test-concourse-resource-template.sh

echo " "

if [ "$1" = "-debug" ]
then
    echo "test-concourse-resource-template.sh -debug (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status. Needed for concourse.
    # set -x enables a mode of the shell where all executed commands are printed to the terminal.
    set -e -x
    echo " "
else
    echo "test-concourse-resource-template.sh (START)"
    # set -e causes the shell to exit if any subcommand or pipeline returns a non-zero status.  Needed for concourse.
    set -e
    echo " "
fi

echo "GOAL ----------------------------------------------------------------------------------"
echo " "

echo "You used your resoruce to get seomthing"
echo "Not open the fetch.json file you put here"
echo " "

echo "CHECK THINGS --------------------------------------------------------------------------"
echo " "

echo "At start, you should be in a /tmp/build/xxxxx directory with two folders:"
echo "   /concourse-resource-template"
# echo "   /concourse-resource-template-updated (created in task-build-push.yml task file)"
echo " "

echo "pwd is: $PWD"
echo " "

echo "List whats in the current directory"
ls -la
echo " "

echo "DO SOMETHING -------------------------------------------------------------------------"
echo " "

# CD WORKING DIRECTORY
WORKING_DIR=$1
echo "WORKING_DIR = $WORKING_DIR"
cd "$WORKING_DIR"
echo "pwd is"
pwd 
echo ""
echo "ls directory"
ls -lat 
echo ""

echo "test-concourse-resource-template.sh (END)"
echo " "
