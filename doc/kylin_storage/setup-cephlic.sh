#!/bin/bash

#部署osd授权文件
path="./ceph.lic"
if [ -f $path ];then
    setup_cephlic=`cat $path | docker exec -i ceph-cmd licence upload`
else
    echo "没有授权文件"
    exit 1
fi
echo ok
