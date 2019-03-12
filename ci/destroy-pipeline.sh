#!/bin/bash
# resource-template destroy-pipeline.sh

fly -t ci destroy-pipeline --pipeline resource-template
