#!/bin/bash

#说明
show_usage="args:  [-h, -d]\
                  [--host= , --disks=]"
#参数
host=""
disks=""
GETOPT_ARGS=`getopt -o h:d -al host:,disks: -- "$@"`
eval set -- "$GETOPT_ARGS"
#获取参数
while [ -n "$1" ]
do
    case "$1" in
        -h|--host) host=$2; shift 2;;
        -d|--disks) disks=$2; shift 2;;
        --) break ;;
        *) echo $1,$2,$show_usage; break ;;
    esac
done

old_disks=$disks
ifs_old=$IFS
IFS=$','
disks=($disks)
IFS="$ifs_old"

#检查网盘可用
for d in ${disks[@]}
do
    ssh $host "mkfs.xfs ${d} -f" > /dev/null 2>&1
    if [ $? -eq '0' ];then
      echo "    $host节点格式化盘$d成功！"
    else
      echo "    $host节点格式化盘$d失败！请检查节点$host网络或$d盘！"
      exit 1
    fi
    docker exec ceph-cmd daemon-create host=$host osd disk=$d
done

echo "osd节点$host部署完成"
    

    










