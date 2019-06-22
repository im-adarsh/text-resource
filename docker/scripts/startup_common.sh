#!/bin/bash
CODEDIR="/go/src/github.com/im-adarsh/text-resource"
echo "setting up config service"
bash /opt/config-service/startup.sh

#dont start confd
sed -i "s|autostart=true|autostart=false|" /etc/supervisor/conf.d/confd.conf

#fix supervisor for https://github.com/docker/docker/issues/12080
sed -i 's|serverurl=unix|serverurl=http://127.0.0.1:9001\n\n[inet_http_server]\nport = 127.0.0.1:9001 \n;|' /etc/supervisor/supervisord.conf

export HOST=`hostname`
sed -i "s/HOSTNAME_HERE/$HOST" /etc/supervisor/conf.d/text-resource.conf

#if consul ip is set, use that instead of local consul
if [ ! -z "$CONSUL_IP" ]
then
    sed -i "s/127.0.0.1/$CONSUL_IP/" $CODEDIR/docker/consul-template/consul-template.conf
elif [ ! -z "$DISABLE_CONSUL" ]
then
    # case where consul is disabled and we are not passed any IP
    IP=`curl "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip" -H "Metadata-Flavor: Google"`
    if [[ -z "$IP" ]]
    then
        IP=`hostname -I`
    fi
    sed -i "s/127.0.0.1/$IP/" $CODEDIR/docker/consul-template/consul-template.conf
fi

if [ ! -z "$CONSUL_PORT" ]
then
    sed -i "s/8500/$CONSUL_PORT/" $CODEDIR/docker/consul-template/consul-template.conf
fi

if [ -z "$HOST_ZONE" ]
then
    p=`curl "http://metadata.google.internal/computeMetadata/v1/instance/zone" -H "Metadata-Flavor: Google"`
    if [ $? -eq 0 ]
    then
        export HOST_ZONE=`basename $p`
    else
        export HOST_ZONE="undefined"
    fi
fi
sed -i "s|environment=|environment=HOST_ZONE='$HOST_ZONE'|" /etc/supervisor/conf.d/supervisor-consul-template.conf
