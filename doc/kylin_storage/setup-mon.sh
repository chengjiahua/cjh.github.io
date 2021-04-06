#!/bin/bash
# time: 2020.2.29
# author: wangfengsheng
# update: 2020.3.2

#说明
show_usage="args:  [-p , -c , -h]\
				  [--public_network= , --cluster_network= , --hosts=]"
#参数
public_network=""
cluster_network=""
hosts=""
GETOPT_ARGS=`getopt -o p:c:h -al public_network:,cluster_network:,hosts: -- "$@"`
eval set -- "$GETOPT_ARGS"
#获取参数
while [ -n "$1" ]
do
	case "$1" in
		-p|--public_network) public_network=$2; shift 2;;
		-c|--cluster_network) cluster_network=$2; shift 2;;
		-h|--hosts) hosts=$2; shift 2;;
		--) break ;;
		*) echo $1,$2,$show_usage; break ;;
	esac
done

old_hosts=$hosts
ifs_old=$IFS
IFS=$','
hosts=($hosts)
IFS="$ifs_old"


echo "+++++++++++++++++++正在分发镜像+++++++++++++++++++"
docker exec ceph-cmd cluster-setup-cmd host="*"                                                                                                                                                                      
echo "+++++++++++++++++管理工具部署成功+++++++++++++++++"


#生成并修改配置文件
gen="docker exec ceph-cmd gen "${hosts[@]}
setup_gen=(`bash -c "$gen"`)
ceph_conf=/etc/ceph/ceph.conf

if [ -f $ceph_conf ];then
	if [[ -n $public_network && -n $cluster_network ]]; then
        sed -i "/^public_network/cpublic_network = $public_network" $ceph_conf
        sed -i "/^cluster_network/ccluster_network = $cluster_network" $ceph_conf
        sed -i "/^management_network/cmanagement_network = $public_network" $ceph_conf
    fi
    sed -i "/^ntp_wait/cntp_wait = false" $ceph_conf
fi

#分发配置文件
sync_ceph_conf=(`bash -c "docker exec ceph-cmd cluser-etc-sync hosts="*""`)

#部署mon节点
for n in ${hosts[@]}
do
    echo "正在部署mon.$n"
    setup_mon=(`bash -c "docker exec ceph-cmd daemon-create host=$n mon"`)
    echo "正在部署mgr.$n"
    setup_mgr=`docker exec ceph-cmd daemon-create host=$n mgr`
done

echo "ok"
