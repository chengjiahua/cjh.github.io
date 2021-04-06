#!/bin/bash

#说明
show_usage="args:  [-t, -i, -p, -d]\
                  [--ctr= , --ip , --cpt= , --disks=]"
#参数
ctr=""
ip=""
cpt=""
disks=""
ceph_conf=/etc/ceph/ceph.conf
GETOPT_ARGS=`getopt -o t:p:d -al ctr:,ip:,cpt:,disks: -- "$@"`
eval set -- "$GETOPT_ARGS"
#获取参数
while [ -n "$1" ]
do
    case "$1" in
        -t|--ctr) ctr=$2; shift 2;;
        -i|--ip) ip=$2; shift 2;;
        -p|--cpt) cpt=$2; shift 2;;
        -d|--disks) disks=$2; shift 2;;
        --) break ;;
        *) echo $1,$2,$show_usage; break ;;
    esac
done


old_ctr=$ctr
old_ip=$ip
old_disks=$disks
ifs_old=$IFS
IFS=$','
ctr=($ctr)
ip=($ip)
disks=($disks)
IFS="$ifs_old"


if [ -f $ceph_conf ]; then
  #添加控制节点
  if [[ -n $ctr && -n $ip ]]; then
    #添加mon_hostname
    for n in ${ctr[@]}; do
       docker exec ceph-cmd cluster-setup-cmd host=$n
       sed -i 's/^mon_initial_members.*/&, '"$n"'/g' $ceph_conf
    done
    #添加mon_ip
    for p in ${ip[@]}; do
       sed -i 's/^mon_host.*/&, '"$p"'/g' $ceph_conf
    done
    docker exec ceph-cmd cluster-etc-sync host=*    
    #推送配置文件，创建mon,mgr
    for n in ${ctr[@]}; do
       #docker exec ceph-cmd cluster-etc-sync host=$n
       docker exec ceph-cmd daemon-create  host=$n mon
       docker exec ceph-cmd daemon-create  host=$n mgr
    done
    #docker exec ceph-cmd cluster-etc-sync host=*
  elif [[ -n $cpt && -n $disks ]];then
  #添加计算节点
  #分发镜像和配置文件
    echo "add cpt"
    echo $cpt $disks
    docker exec ceph-cmd cluster-setup-cmd host=$cpt
    docker exec ceph-cmd cluster-etc-sync host=$cpt
    #检查网盘可用
    for d in ${disks[@]};
    do
        ssh $cpt "mkfs.xfs ${d} -f" > /dev/null 2>&1
        if [ $? -eq '0' ];then
          echo "    $cpt节点格式化盘$d成功！"
        else
          echo "    $cpt节点格式化盘$d失败！请检查节点$cpt网络或$d盘！"
          exit 1
        fi
        docker exec ceph-cmd daemon-create host=$cpt osd disk=$d
    done
    echo "osd节点$cpt部署完成"
  else
    #只加节点，不部署osd
    docker exec ceph-cmd cluster-setup-cmd host=$cpt
    docker exec ceph-cmd cluster-etc-sync host=$cpt
  fi
fi


