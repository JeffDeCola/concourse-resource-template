#!/bin/bash
# resource-template set-pipeline.sh

fly -t ci set-pipeline -p resource-template -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
