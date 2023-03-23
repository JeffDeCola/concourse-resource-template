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

echo "To do something to the STUFF you got"
echo " "

echo "CHECK THINGS --------------------------------------------------------------------------"
echo " "

echo "At start, you should be in a /tmp/build/xxxxx directory with two folders:"
echo "   /concourse-resource-template"
# echo "   /concourse-resource-template-updated (created in task-build-push.yml task file)"
echo " "

echo "pwd is"
pwd 
echo ""

echo "ls directory"
ls -lat 
echo ""

echo "DO SOMETHING -------------------------------------------------------------------------"
echo " "

echo "cd jeffs-test-resource"
cd jeffs-test-resource
echo "List whats in the current directory"
ls -la
echo " "

echo "cat get_fetch.json"
cat get_fetch.json
echo " "

echo "Add a file for fun"
echo 'This is a test' > job-test-concourse-resource-template.txt
echo " "

echo "ls directory"
ls -lat 
echo ""

echo "test-concourse-resource-template.sh (END)"
echo " "
