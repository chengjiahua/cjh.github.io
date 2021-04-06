ceph_cmd=`docker ps | grep ceph-cmd | awk '{printf $1}'`
echo $ceph_cmd
if [ -n $ceph_cmd ];then
  pool_list=`docker exec ceph-cmd ceph osd pool ls`
  if [ -n "$pool_list" ];then
    `docker exec ceph-cmd echo mon_allow_pool_delete = true >> /etc/ceph/ceph.conf`
    echo "mon_allow_pool_delete参数修改完成"
    `docker exec ceph-cmd daemon-restart host=* target=mon`
    echo "mon服务重启完成"
    for p in $pool_list
    do
        `docker exec ceph-cmd ceph osd pool delete $p $p --yes-i-really-really-mean-it`
        echo "pool $p 删除完成"
    done
  else
     echo "没有pool资源"
  fi
else
  echo "ceph-cmd工具不存在"
fi
