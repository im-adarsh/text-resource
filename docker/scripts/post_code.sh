#!/bin/bash
CODEDIR='/go/src/github.com/im-adarsh/text-resource'
cp $CODEDIR/docker/supervisor-consul-template.conf /etc/supervisor/conf.d/
cp $CODEDIR/docker/text-resource.conf /etc/supervisor/conf.d/
cp $CODEDIR/docker/commands/* /usr/bin/
