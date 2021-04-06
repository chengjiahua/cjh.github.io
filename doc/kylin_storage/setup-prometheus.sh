#!/bin/bash

#说明
show_usage="args:  [-p, -g]\
                  [--prometheus_ip= , --grafana_ip=]"
#参数
prometheus_ip=""
grafana_ip=""
GETOPT_ARGS=`getopt -o p:g -al prometheus_ip:,grafana_ip: -- "$@"`
eval set -- "$GETOPT_ARGS"
#获取参数
while [ -n "$1" ]
do
	case "$1" in
		-p|--prometheus_ip) prometheus_ip=$2; shift 2;;
		-g|--grafana_ip) grafana_ip=$2; shift 2;;
		--) break ;;
		*) echo $1,$2,$show_usage; break ;;
	esac
done
 
	echo "正在部署prometheus"
	docker exec ceph-cmd daemon-create host=$prometheus_ip prometheus
	echo "正在部署grafana"
	docker exec ceph-cmd daemon-create host=$grafana_ip grafana
