#!/bin/bash
# jeffCoin set-pipeline.sh

fly -t ci set-pipeline -p jeffCoin -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
