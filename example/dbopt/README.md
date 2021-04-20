# module-demo

modeule 服务名字
demo 服务中的子模块

## 项目简介

modeule-demo demo演示程序

## 环境安装

1. 编译

    1. 下载1.13以上的go安装包
      
    wget https://studygolang.com/dl/golang/go1.13.14.linux-arm64.tar.gz

    2. 解压go安装包
    tar xvf go1.13.14.linux-arm64.tar.gz -C /usr/local/

    3. 修改环境变量
    vim /etc/profile 在里面添加
    export GOROOT=/usr/local/go 
    export GOPATH=$HOME/golang 
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    export GOPROXY=https://goproxy.cn

    4. source /etc/profile
    go version
    chmod +x control

    5. 安装protobuf
    从172.20.42.1 x86物理机上的/root/google 拷贝你自己环境上的golang/src/目录下

    6. 生成probuf文件
    kops目录下执行./control proto module-demo
    执行后kops/app/module-demo/api/v1目录下module-demo.bm.go  module-demo.ecode.go module-demo.pb.go module-demo.swagger.json四个文件就由api.proto文件自动生成了
    
    7. 编译运行模块
    kops目录下执行./control run module-demo
    kops/bin目录下会生成一个module-demo目录，该目录下module-demo即为二进制文件 

    8. 自动更新 Swagger 接口数据到 YApi 平台
    安装yapi-cli工具yum install npm ; npm install yapi-cli [-g] (确保 yapi-cli 工具版本 >= 1.2.7)
    如果安装失败，则是yapi-cli 依赖 Node.js , 请安装不低于 7.6 版本的 Node.js 
    在kops/app/module-demo/api/v1创建yapi-import.json文件

    yapi-import.json文件格式为
    {
      "type": "swagger",
      "token": "eb1727efd30adaa4990d0f8f178e731bbf94e55aab05a57d9d10cc048846af2e",
      "file": "module-demo.swagger.json",
      "merge": "mergin",
      "server": "http://api.kylincloud.me/"
    }

    type 是数据数据方式，目前官方只支持 swagger
    token为项目唯一指定的token，在 项目设置 -> token 设置获取
    file 是 swagger 接口文档文件，protobuf文件中的module-demo.swagger.json文件
    merge 导入旧的接口策略，默认使用智能模式，一共有 "normal"(普通模式) , "good"(智能合并), "merge"(完全覆盖) 三种模式
    server 是yapi服务器地址
    官网YApi教程网址：https://hellosean1025.github.io/yapi/documents/plugin-list.html

    配置完成后，kops目录下执行./control proto module-demo import
    即可自动生成protobuf文件并自动上传到 YApi 平台，更新 Swagger 接口数据
    
2. 数据库

    * 安装sqlite3数据库:

3. 运行

    ```
    在kops目录下执行
    chmod +x control
    ./control run module-demo
    即可编译运行module-demo 模块
    ```


