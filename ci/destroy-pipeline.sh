#!/bin/sh
# concourse-resource-template destroy-pipeline.sh

echo " "
echo "Destroy pipeline on target jeffs-ci-target which is team jeffs-ci-team"
fly --target jeffs-ci-target \
    destroy-pipeline \
    --pipeline concourse-resource-template
echo " "
