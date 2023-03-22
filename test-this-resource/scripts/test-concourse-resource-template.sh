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

echo "cd jeffs-test-resource"
cd jeffs-test-resource
echo "List whats in the current directory"
ls -la
echo " "

echo "test-concourse-resource-template.sh (END)"
echo " "
