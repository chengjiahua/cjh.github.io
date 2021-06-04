#!/bin/bash

#修改docker0和docker源
systemctl stop docker

ip link set dev docker0 down

filename="/etc/docker/daemon.json"

if [ ! -f ${filename} ];then
    cp package/docker/daemon.json /etc/docker/daemon.json 
else
    echo "${filename},此文件已存在,请手动更改"
    exit -1
fi

 systemctl restart docker