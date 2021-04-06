#!/bin/bash

#获取osd的数量
osd_num=`docker exec  ceph-cmd daemon-list host=* target=osd | grep osd | wc -l`
#获取mon的数量
mon_num=`docker exec  ceph-cmd daemon-list host=* target=mon | grep mon | wc -l`
#设置pg_num
pg_num=$[$osd_num*100/3/4]
flag=1
until [[ "$flag" -gt $pg_num ]]
do
   flag=$[$flag*2]
done
pg_num=$[$flag/2]
pool_lists="volumes images vms backups"
old_pg_num=`docker exec ceph-cmd ceph osd pool get volumes pg_num | awk '{print $2}'`
if [ $pg_num -gt $old_pg_num ];then
  for p in $pool_lists
  do
    docker exec ceph-cmd ceph osd pool set $p pg_num $pg_num
    docker exec ceph-cmd ceph osd pool set $p pgp_num $pg_num
  done
fi
