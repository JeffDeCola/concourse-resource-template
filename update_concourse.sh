#!/bin/bash
# resource-template update_concourse.sh

fly -t ci set-pipeline -p resource-template -c ci/pipeline.yml --load-vars-from ../../../../.credentials.yml
