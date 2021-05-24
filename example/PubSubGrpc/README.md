grpc实现订阅推送模式简单实例

利用grpc 服务器流的交互模式与docker社区的pubsub包实现

grpc-service服务起来后，client1向service端推送消息（调用grpc流），然后client2监听到消息（需要轮询），完成订阅过程。

实例用法：

grpc-service服务起来
cd ./service; go run main

client2监听到消息（需要轮询）
cd ./client2; go run client2.go

client1向service端推送消息（调用grpc流）
cd ./client; go run client.go