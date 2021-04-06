    1.下载1.13以上的go安装包  
    wget https://studygolang.com/dl/golang/go1.13.14.linux-arm64.tar.gz
    2.解压go安装包
    tar xvf go1.13.14.linux-arm64.tar.gz -C /usr/local/

    3.vim /etc/profile 在里面添加
    export GOROOT=/usr/local/go 
    export GOPATH=$HOME/golang 
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    export GOPROXY=https://goproxy.cn

    4.source /etc/profile
    go version查看当前版本号是否一直

    5.编译模块
    编译单个模块: `./control build kops-api` 
    编译全部模块: `./control build ` 如有新增模板需修改control脚本
    kops目录下执行./control build kops-api
    kops/bin目录下会生成一个kops-api目录，该目录下kops-api即为二进制文件

    6.安装protobuf
    从83.65物理机上的/data/qcow2s拷贝protoc-3.12.4目录
    mkdir -p $GOPATH/bin
    cp -r protoc-3.12.4/bin/*  $GOPATH/bin/
    cp -r protoc-3.12.4/include/google $GOPATH/src/

    7.生成probuf文件
    kops目录下执行./control proto kops-api
    成功执行后kops/app/kops-api/api/v1目录下api.bm.go  api.ecode.go api.pb.go api.swagger.json四个文件就由api.proto文件自动生成了

    sed -i "s/kops-demo/monitor-judge/g" `grep kops-demo -rl monitor-judge/`