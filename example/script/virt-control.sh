#!/bin/bash

CWD=$(cd $(dirname $0)/; pwd)
cd $CWD

os_qcow2_name=""
disk_qcow2_name=""

usage()
{
	echo -e "Flag shorthand -h has been deprecated, please use --help\n"
	echo -e "network control bash version = 0.1.0\n"
	echo -e "Usage: bash $0  {creat} \n"
	echo -e "Usage Module:\n"
    echo -e "创建目标节点 (即磁盘都是空的)"
	echo -e "	creat <name> <disk_dir> <cpu_num(4)>  <memory_size(8G)>  <eth0_bridge(br0)>  <eth1_bridge(br0)> \n"


	exit 0
}

creat()
{
    check_format $1 $2 $3 $4 $5 $6
    if [ $? -ne 0 ];then
        echo "创建虚拟机失败" 
        return 1
    fi
}

check_format()
{
    name=$1
    disk_dir=$2
    cpu_num=$3
    memory_size=$4
    eth0_bridge=$5
    eth1_bridge=$6

    if [ "x${name}" = "x" ]; then
        echo "虚拟机名称不能为空" 
        return 1
    fi

    # 判断该虚拟机是否已存在
    is_exist=`virsh list --all | grep ${name} `
    if [ "x${is_exist}" != "x" ]; then
        echo "该虚拟机已存在" 
        return 1
    fi

    if [ "x${disk_dir}" = "x" ]; then
        echo "磁盘目录路径不能为空" 
        return 1
    else
        if [ ! -d "${disk_dir}" ]; then
        echo "磁盘目录路径不存在" 
        return 1
        fi
    fi


    if [ "x${cpu_num}" = "x" ]; then
        cpu_num=4
    fi

    if [ "x${memory_size}" = "x" ]; then
        memory_size=8388608
    else
        temp=$((memory_size*1024*1024))
        memory_size=$temp
    fi

    if [ "x${eth0_bridge}" = "x" ]; then
        eth0_bridge=br0
    fi

    if [ "x${eth1_bridge}" = "x" ]; then
        eth1_bridge=br0
    fi

    echo "确认添加虚拟机参数"
    echo "虚拟机名称：$name"
    echo "内存大小：$((memory_size/1024/1024))G  cpu核数: $cpu_num"核
    echo "磁盘目录路径：$disk_dir "
    echo "eth0网桥：${eth0_bridge}, eth1网桥：${eth1_bridge}"

    creat_start $name $disk_dir $cpu_num $memory_size $eth0_bridge $eth1_bridge
    if [ $? -ne 0 ];then
        return 1
    fi
}

creat_start()
{
    name=$1
    disk_dir=$2
    cpu_num=$3
    memory_size=$4
    eth0_bridge=$5
    eth1_bridge=$6

    # 判断架构，拷贝xml模板
    prexml $name
    if [ $? -ne 0 ];then
        echo "拷贝xml模板文件失败"
        return 1
    fi

    # 创建磁盘
    creat_disk $disk_dir
    if [ $? -ne 0 ];then
        echo "创建磁盘失败"
        return 1
    fi

    # 更新xml配置文件
    update_xml $name $disk_dir $cpu_num $memory_size $eth0_bridge $eth1_bridge
    if [ $? -ne 0 ];then
        echo "更新xml配置文件失败"
        return 1
    fi

    # 创建虚拟机
    virsh define $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "创建虚拟机失败"
        return 1
    fi

    # 启动虚拟机
    virsh start ${name}
    if [ $? -ne 0 ];then
        echo "启动虚拟机失败"
        return 1
    fi

    # 增加网卡
    virsh attach-interface ${name} --type  bridge  --source ${eth0_bridge} --model virtio  --persistent
    if [ $? -ne 0 ];then
        echo "增加网卡eth0失败"
        return 1
    fi

    virsh attach-interface ${name} --type  bridge  --source ${eth1_bridge} --model virtio  --persistent
    if [ $? -ne 0 ];then
        echo "增加网卡eth1失败"
        return 1
    fi
}

prexml()
{

    hw_arch=$(uname -m)

    x86_archs=(
        "amd64"
        "x86_64"
        "x86"
    )

    arm_archs=(
        "aarch64"
        "arm64"
    )


    for arch in ${arm_archs[@]}; do
        if [ "x${arch}" = "x${hw_arch}" ]; then
            xml="package/xml/arm-master.xml"
        fi
    done

    for arch in ${x86_archs[@]}; do
        if [ "x${arch}" = "x${hw_arch}" ]; then
            xml="package/xml/x86-master.xml"
        fi
    done
    
    \cp $CWD/$xml $CWD/${name}.xml
    if [ $? -ne 0 ];then
        return 1
    fi

}

update_xml()
{
    name=$1
    disk_dir=$2
    cpu_num=$3
    memory_size=$4
    eth0_bridge=$5
    eth1_bridge=$6

    sed -i "s/NAMEMODE/$name/g" $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "更新xml文件虚拟机名称失败"
        return 1
    fi

    sed -i "s/MEMORYSIZE/$memory_size/g" $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "更新xml文件内存大小失败"
        return 1
    fi

    sed -i "s/CPUNUM/$cpu_num/g" $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "更新xml文件cpu核数失败"
        return 1
    fi
    
    # 更新磁盘信息
    sed -i "s:QCOW2PATH:${disk_dir}/${os_qcow2_name}:g" $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "更新os_disk失败"
        return 1
    fi

    sed -i "s:QCOW2DISKPATH:${disk_dir}/${disk_qcow2_name}:g" $CWD/${name}.xml
    if [ $? -ne 0 ];then
        echo "更新data_disk失败"
        return 1
    fi
}

creat_disk()
{
    disk_dir=$1
    i=1
    while [ $i -le 20 ]
    do
        check_disk $disk_dir $i
        if [ $? -ne 0 ];then
            # echo "文件已存在"
            i=$((i+1))
            continue
        else
            cd $disk_dir;qemu-img create -f qcow2 master${i}.qcow2 300G
            if [ $? -ne 0 ];then
                echo "master${i}.qcow2 创建失败"
                return 1
            fi
            cd $disk_dir;qemu-img create -f qcow2 master${i}-disk.qcow2 300G
            if [ $? -ne 0 ];then
                echo "master${i}-disk.qcow2 创建失败"
                return 1
            fi

            os_qcow2_name="master${i}.qcow2"
            disk_qcow2_name="master${i}-disk.qcow2"

            break
        fi
    done
}

check_disk()
{
    disk_dir=$1
    i=$2
    for file_a in ${disk_dir}/*
    do  
        file=`basename $file_a`  
            if [ "x${file}" = "xmaster${i}.qcow2" ]; then 
                return 1
            else
                continue
            fi  
    done
}



case "$1" in
	creat)
		creat $2 $3 $4 $5 $6 $7
		;;
	*)
		usage
esac