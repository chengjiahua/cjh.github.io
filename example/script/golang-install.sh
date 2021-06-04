#!/bin/bash
# go环境初始化脚本，默认go语言压缩包在/root目录，且只有一个压缩包,默认架构正确

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

# 确定架构

  for arch in ${arm_archs[@]}; do
    if [ "x${arch}" = "x${hw_arch}" ]; then
        package_name="package/go/go1.14.13.linux-arm64.tar.gz"
    fi
  done

  for arch in ${x86_archs[@]}; do
    if [ "x${arch}" = "x${hw_arch}" ]; then
        package_name="package/go/go1.14.10.linux-amd64.tar.gz"
    fi
  done


# 删除以前的go环境
rm -rf /usr/local/go

# 解压go安装包
tar xvf $package_name -C /usr/local/ 1>/dev/null

if [ $? -ne 0 ]; then
    echo "解压go语言包失败"
    exit -1
fi

# 修改环境变量
    export GOROOT=/usr/local/go 
    export GOPATH=$HOME/golang 
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    export GOPROXY=https://goproxy.cn

echo "
    export GOROOT=/usr/local/go 
    export GOPATH=\$HOME/golang 
    export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin
    export GOPROXY=https://goproxy.cn
" >> /etc/profile

source /etc/profile

if [ $? -ne 0 ]; then
    echo "修改环境变量失败"
    exit -2
fi

if [ "x$(go version)" == "x-bash: go" ]; then 
    echo "go环境安装失败"
    exit -3
fi