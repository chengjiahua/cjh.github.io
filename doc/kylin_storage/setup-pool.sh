#!/bin/bash

#获取osd的数量
osd_num=`docker exec  ceph-cmd daemon-list host=* target=osd | grep osd | wc -l`
#获取mon的数量
mon_num=`docker exec  ceph-cmd daemon-list host=* target=mon | grep mon | wc -l`
#设置pg_num
pg_num=$[$osd_num*200/3/4]
flag=1
until [[ "$flag" -gt $pg_num ]]
do
   flag=$[$flag*2]
done
pg_num=$[$flag/2]
if [ $pg_num -lt 32 ];then
  pg_num=32
fi
#设置资源池
pool_lists="volumes images vms backups"
for p in $pool_lists
do 
   set_pool=`docker exec ceph-cmd ceph osd pool create $p $pg_num`
   echo "pool $p 设置成功"
done

 \cp -r /kylincloud/config/ /etc/kolla/                                            
 \cp -rf /etc/ceph/*   /etc/kolla/config/cinder/
 mkdir -p /etc/kolla/config/cinder/cinder-backup
 mkdir -p /etc/kolla/config/cinder/cinder-volume
 \cp /etc/ceph/ceph.client.admin.keyring  /etc/kolla/config/cinder/cinder-backup
 \cp /etc/ceph/ceph.client.admin.keyring  /etc/kolla/config/cinder/cinder-volume
 \cp -rf /etc/ceph/*   /etc/kolla/config/glance/
 \cp -rf /etc/ceph/*   /etc/kolla/config/nova/
 echo "文件目录配置成功" 
