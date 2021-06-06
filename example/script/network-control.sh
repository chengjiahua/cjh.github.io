#!/bin/bash

usage()
{
	echo -e "Flag shorthand -h has been deprecated, please use --help\n"
	echo -e "network control bash version = 0.1.0\n"
	echo -e "Usage: bash $0  {update} \n"
	echo -e "Usage Module:\n"
	echo -e "	update <interface>  <ip>  <gw>  <mask>    			update network information"
	exit 0
}

update()
{
	interface=$1
	ip=$2
	gw=$3
	mask=$4

# 检查参数格式
	check_interface $interface
	if [ $? -ne 0 ]
    then
        echo "$interface网卡不存在" 
        return 1
    fi
	
	check_ipaddr $ip
	if [ $? -ne 0 ]
    then
        echo "ip地址必须全部为数字" 
        return 1
    fi

	check_ipaddr $gw
	if [ $? -ne 0 ]
    then
        echo "gw地址必须全部为数字" 
        return 1
    fi

	check_ipaddr $mask
	if [ $? -ne 0 ]
    then
        echo "mask地址必须全部为数字" 
        return 1
    fi

# 开始修改网卡信息
	update_start $interface $ip $gw $mask
	if [ $? -ne 0 ]
    then
        echo "修改网卡信息失败" 
        return 1
    fi
}

check_interface()
{
	if [ "x$(ls /etc/sysconfig/network-scripts/ | grep $1)" = "x" ];then
		return 1
	fi
}

check_ipaddr()
{
    echo $1|grep "^[0-9]\{1,3\}\.\([0-9]\{1,3\}\.\)\{2\}[0-9]\{1,3\}$" > /dev/null;
    if [ $? -ne 0 ];then
        return 1
    fi
    ipaddr=$1
    a=`echo $ipaddr|awk -F . '{print $1}'`  #以"."分隔，取出每个列的值 
    b=`echo $ipaddr|awk -F . '{print $2}'`
    c=`echo $ipaddr|awk -F . '{print $3}'`
    d=`echo $ipaddr|awk -F . '{print $4}'`
    for num in $a $b $c $d
    do
        if [ $num -gt 255 ] || [ $num -lt 0 ]    #每个数值必须在0-255之间 
        then
            echo $ipaddr "中，字段"$num"错误" 
            return 1
        fi
   done
   return 0
}

update_start()
{
	if [ "x${interface}" = "x" ]; then
		echo "网卡名不能为空"
		return -1
	else
		if [ "x${ip}" = "x" ]; then
			echo "ip不能为空"
			return -2
		else
			# 修改ip
			sed -i "/IPADDR/c\IPADDR=$ip" /etc/sysconfig/network-scripts/ifcfg-${interface}
			if [ $? -ne 0 ]; then
				echo "修改ip失败"
				return -3
			fi


			if [ "x${gw}" != "x" ]; then
				if [ "x${mask}" != "x" ]; then
					#修改网关+mask
					sed -i "/GATEWAY/c\GATEWAY=${gw}" /etc/sysconfig/network-scripts/ifcfg-${interface}
					if [ $? -ne 0 ]; then
						echo "修改网关失败"
						return -3
					fi

					sed -i "/NETMASK/c\NETMASK=${mask}" /etc/sysconfig/network-scripts/ifcfg-${interface}
					if [ $? -ne 0 ]; then
						echo "修改掩码失败"
						return -3
					fi

				else
					#只修改网关
					sed -i "/GATEWAY/c\GATEWAY=${gw}" /etc/sysconfig/network-scripts/ifcfg-${interface}
					if [ $? -ne 0 ]; then
						echo "修改网关失败"
						return -3
					fi
				fi
			fi

			#重启网络
			ifdown ${interface}
			ifup ${interface}
		fi
	fi
}

case "$1" in
	update)
		update $2 $3 $4 $5
		;;
	*)
		usage
esac