#!/bin/bash

ceph_conf="/etc/ceph/ceph.conf"
if [ -f $ceph_conf ];then
  #删除osd服务
  delosd=`docker exec ceph-cmd daemon-delete host=* target=osd`
  echo "osd服务删除完成"
  #删除mgr服务
  delmgr=`docker exec ceph-cmd daemon-delete host=* target=mgr`
  echo "mgr服务删除完成"
  #删除mon服务
  delmmon=`docker exec ceph-cmd daemon-delete host=* target=mon`
  echo "mon服务删除完成"
  rm -rf /etc/ceph/ceph*
else
  echo "已经删除完毕"
fi
