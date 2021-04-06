部署的前提条件是：
1.ceph的镜像已经在控制节点主机上
2.各节点的/etc/hosts和/etc/ceph/hosts配置完成
3.各主机互相免密

部署流程:
1.运行setup-mon脚本. 例如: ./setup-mon --public_network=1.2.3.0/24 --cluster_network=1.2.3.0/24 --hosts=controller,compute1,compute2
2.运行setup-cephlic脚本. 例如：./setup-cephlic(在部署osd服务之前执行)
3.运行setup-osd脚本. 例如: ./setup-osd --host=controller --disks=/dev/sdb,/dev/sdc 注意: 每次运行该脚本只部署一台主机，可重复运行该脚本。
4.运行check-ceph脚本，例如: ./check-ceph 用来确定ceph集群的状态
