#!/bin/bash
docker build . -t text-resource || exit
docker stop text-resource
docker rm text-resource
docker run -d --privileged --name text-resource -p 9281:9281 -p 9282:9282 -p 9283:9283 -p 9284:9284 -h `whoami`-`hostname` text-resource start_stage
echo ""
echo "Type 'supervisorctl status' to view services status"
echo "Type 'logs' to view app logs"
echo "Type 'code' to go to app code"
echo "Type 'slogs' to wait 20 seconds and then start tailing logs"
docker exec -ti text-resource bash