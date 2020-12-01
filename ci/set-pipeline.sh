#!/bin/bash
# jeffCoin set-pipeline.sh

fly -t ci set-pipeline -p jeffcoin -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
