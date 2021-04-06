#!/bin/bash
NET="$1"
HOST_IP=$(ifconfig $NET  | grep inet | grep -v 127.0.0.1 | grep -v inet6 | awk '{print $2}')
CTR_HOST="$2"
CTR_IP=$(sed -n "/$CTR_HOST/p"  /etc/hosts | awk  '{print $1}')
SERVER_GRAFANA_ID=$(ssh $CTR_HOST  "docker ps | grep grafana " | awk '{print $1}')


`docker exec -d ko_nginx bash -c "echo $CTR_IP grafanaIp >> /etc/hosts"`
`docker exec -d ko_nginx bash -c "nginx -t && nginx -s reload"`

`ssh $CTR_HOST "docker exec -d $SERVER_GRAFANA_ID bash -c  \"sed -i '/^.*root_url /croot_url = http://$HOST_IP:8555/grafana/' /ceph/etc/grafana/grafana.ini\""`
`docker exec -d ceph-cmd bash -c "daemon-restart host=* target=grafana"`
