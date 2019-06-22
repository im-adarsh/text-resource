#!/bin/bash
#config locale
echo -e 'LANG="en_US.UTF-8"\nLANGUAGE="en_US:en"\n' > /etc/default/locale
echo 'en_US.UTF-8 UTF-8' >> /etc/locale.gen
#locale-gen --purge en_US.utf-8

#setup term
echo 'export TERM=screen-256color' >> ~/.bashrc
echo 'tail -f /tmp/carousell*' >> /usr/bin/logs
echo 'sleep 25; logs' >> /usr/bin/slogs
chmod +x /usr/bin/logs
chmod +x /usr/bin/slogs
echo 'alias code="cd /go/src/github.com/im-adarsh/text-resource"' >> ~/.bashrc
