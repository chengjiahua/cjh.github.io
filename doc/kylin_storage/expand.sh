#!/bin/bash
pool_list=`docker exec ceph-cmd ceph osd pool ls`
for i in $pool_list
do 
    echo $i
    set_size=`docker exec ceph-cmd ceph osd pool set $i size 1&&docker exec ceph-cmd ceph osd pool set $i min_size 1`
done
