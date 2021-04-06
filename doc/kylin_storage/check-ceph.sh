#!/bin/bash

ceph_mon=`docker exec ceph-cmd daemon-list host=* target=mon | grep mon`
if [[ -n $ceph_mon ]];then
#    ceph_health=`docker exec ceph-cmd ceph -s | grep health | awk '{print $2}'`
    ceph_health=`docker exec ceph-cmd ceph health | awk '{print $1}'`
    is_clock=`docker exec ceph-cmd ceph health | awk '{print $2}'`
    echo "ceph的状态是$ceph_health"
    if [ $ceph_health != "HEALTH_OK" ];then
       if [ $is_clock != "clock" ];then
         echo "error"
         exit 1
       fi
    else
        echo "OK"
    fi
else
    echo "mon服务没有部署"
    echo "error"
    exit 1
fi
