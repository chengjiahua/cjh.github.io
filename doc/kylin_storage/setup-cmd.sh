#!/bin/bash
#部署管理工具

tag="13.2.5"

if [ "$(uname -m)" == "aarch64" ]; then
  image="registryserver:4000/ft/neokylin-binary-ceph:$tag"
elif  [ "$(uname -m)" == "x86_64" ]; then
  image="registryserver:4000/x86/neokylin-binary-ceph:$tag"
else
  echo "unknown"
  exit 1
fi

#判断是否部署管理工具
is_cmd=`docker ps | grep ceph-cmd | awk '{print $NF}'`
if [ ! $is_cmd ];then
  echo "+++++++++++++++++正在部署管理工具+++++++++++++++++"
  docker_run=(`bash -c "docker run --rm \
    -v /var/run/docker.sock:/var/run/docker.sock \
    $image \
    setup-cmd"`)
  if [ $? -eq '0' ];then
      stat=(`bash -c "docker exec ceph-cmd echo done"`)
      if [ ! $stat ];then
          echo "+++++++++++++++++管理工具部署失败+++++++++++++++++"
          exit 1
      fi
  fi
else
  echo "已经部署了setup-cmd"
fi
