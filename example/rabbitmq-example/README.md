# RabbitMQ的go client实例

用于学习go 是如何与RabbitMQ交互的

## RabbitMQ

RabbitMQ 是一个由 Rabbit 开发的基于 AMQP 的开源消息队列，能够实现异步消息处理，使用 erlang 语言开发。

## 消息发送过程

ConnectionFactory、Connection、Channel 这三个时 RabbitMQ 对外封装的重要对象。Connection 封装了 socket 协议相关逻辑。ConnectionFactory 是 Connection 的制造工厂，Channel 类似于 Java NIO 的 Channel 做选择 TCP 连接复用，不仅可以减少性能开销，同时也便于管理。

## 组件

### 队列（Queue）

队列是用来存储消息的，RabbitMQ 的消息只能存在队列里面，生产者产生消息发送至交换机再然后经过一系列的路由规则最终投递到队列，消费者可以从队列中获取消息。

为了处理消费者没有处理完消息就宕机的情况，我们可以要求消费者消费完消息之后需要发送一个回执给 rabbitmq,rabbitmq 收到回执之后才会将这条消息从队列里面移除。

如果希望 rabbitmq 宕机时不会丢失消息，我们可以将 queue 和 message 都设置为可持久化(durable)。

声明
>- name: 队列名称
>- durable: 消息是否持久化
>- auto-deleted: 队列接受完消息是否自动删除
>- exclusive: 是否为独占队列 (独占队列只能由声明他们的连接访问，并在连接关闭时删除)
>- no-wait: 如果为 true,队列默认认为已存在交换机 Exchange,连接不上会报错
>- argument: 队列的其他选项

```
ch.QueueDeclare(
	r.QueueName,
	false, //是否持久化
	false, //是否为自动删除
	false, //是否具有排他性
	false, //是否阻塞
	nil,   //额外属性
)
```
### 消费(Consume)

>- name: 队列名称
>- consumer: 消费者名称
>- autoAck: 自动应答
>- exclusive: 排他性
>- noLocal: 不允许本地消费(使用同一个 connection)
>- nowait: 是否阻塞
>- args: 其他参数

```
ch.Consume(
	q.Name,
	"",   //用来区分多个消费者
	true, //是否自动应答,告诉我已经消费完了
	false,
	false, //若设置为true,则表示为不能将同一个connection中发送的消息传递给这个connection中的消费者.
	false, //消费队列是否设计阻塞
	nil,
)

```

### 交换机(Exchange)

生产者会将消息发送到 Exchange, 由 Exchange 将消息路由至一个或者多个队列(或者丢弃)。生产者将消息发送给 Exchange 时会指定 routing key 来指定该消息的路由规则。

声明
>- name: 交换机名称
>- type: 交换机类型
>- durable: 是否持久化
>- auto-deleted: 当关联的队列都删除之后自动删除
>- internal: 是否为 rabbitmq 内部使用
>- no-wait: 如果为 false,则不期望 Rabbitmq 服务器有一个 Exchange.DeclareOk 这样的响应
>- argument: 其他选项

```
ch.ExchangeDeclare(
	"eventti.event", // name
	"fanout",        // type
	true,            // durable
	false,           // auto-deleted
	false,           // internal
	false,           // no-wait
	nil,             // arguments
)

```

### 发布(Publish)

```
ch.Publish(
	r.ExChange, // 交换机名称
	"",         // 路由键
	false,      // 消息发送成功确认(没有队列会异常)
	false,      // 消息发送失败回调(队列中没有消费者会异常)
	amqp.Publishing{ // 发送的消息
		ContentType: "text/plain",
		Body:        []byte(message),
	})

```

### 绑定(Binding)

RabbitMQ 中通过 Binding 将 Exchange 和 queue 关联起来，这样就可以正确的路由到对应的队列。

在绑定交换机和队列时通常会指定一个 binding key 当 binding key 和生产者指定的 routing key 相匹配的时候，消息就会被路由到对应的队列中。

binding key 不是一定会生效，要看交换机的类型，比如类型时 fanout，则会进行广播，将消息发送到所有绑定的队列

### 交换机类型(Exchange Types)

Exchange Type 有 fanout 1、direct、topic、headers 这四种。

>- fanout
>- fanout 类型会把所有发送到 fanout Exchange 的消息都会被转发到与该 Exchange 绑定(Binding)的所有 Queue 上。就是广播。
>- direct
>- direct 类型会把消息路由到那些 binding key 与 routing key 完全匹配的 Queue 中。
>- topic
>- topic 类型在 direct 类型的匹配规则上有约束：
>- routing key 是一个句点号"."分隔的字符串
>- binding key 也是一个句点号"."分隔的字符串
>- binding key 中存在两种特殊字符 *、# 进行模糊匹配,其中 * 匹配一个单词，# 匹配零个或者多个单词。实例：a.b.c 会被匹配到 *.b.* 和 *.*.c
>- headers
>- 消息发布前,为消息定义一个或多个键值对的消息头,然后消费者接收消息同时需要定义类似的键值对请求头:(如:x-mactch=all 或者 x_match=any)，只有请求头与消息头匹配,才能接收消息,忽略 RoutingKey。

## RabbitMQ 的四种模式

### 简单(点对点)模式
点对点,一个生产者产生消息发送至消息队列，一个消费者消费。

### 订阅(发布)模式
只要绑定了当前交换机的队列就能收到消息。

### 路由(routing)模式
根据 routing key 和 binding key 完全匹配的路由规则进行分发。

### 主题(模糊匹配)模式
也就是 topic 类型的交换类型,与路由模式相比，可进行模糊匹配，如果 Exchange 没有发现能与 routing key 匹配的队列，则会丢弃消息。