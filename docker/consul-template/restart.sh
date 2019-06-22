#!/bin/bash
echo "reloading config"
pid=$(supervisorctl pid im-adarsh)
if [[ $pid -eq 0 ]]
then
    supervisorctl start im-adarsh
else
    supervisorctl pid im-adarsh| xargs kill -s SIGHUP
fi
