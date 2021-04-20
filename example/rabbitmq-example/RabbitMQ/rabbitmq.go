package RabbitMQ

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const MQURL = "amqp://guest:guest@127.0.0.1:5672"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！")

	return rabbitmq
}

// Destory 断开channel和connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQSimple
// 简单模式Step 1.创建简单模式下的RabbitMq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式Step 2:简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message interface{}) {
	byteMsg, err := json.Marshal(message)
	r.failOnErr(err, "JSON Marshal failed")
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err = r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        byteMsg,
		},
	)
}

// ConsumeSimple 使用 goroutine 消费消息
func (r *RabbitMQ) ConsumeSimple() {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协和处理消息
	go func() {
		for d := range msgs {
			// 实现我们要实现的逻辑函数
			log.Printf("Received a message: %s", string(d.Body))
			// fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}

//这里是订阅模式的相关代码。
//订阅模式需要用到exchange。
//获取订阅模式下的rabbitmq的实例
func NewRabbitMqSubscription(exchangeName string) *RabbitMQ {
	//创建rabbitmq实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "订阅模式连接rabbitmq失败。")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "订阅模式获取channel失败")
	return rabbitmq
}

//订阅模式发布消息
func (r *RabbitMQ) PublishSubscription(message string) {
	//第一步，尝试连接交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", //这里一定要设计为"fanout"也就是广播类型。
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式发布方法中尝试连接交换机失败。")

	//第二步，发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//订阅模式消费者
func (r *RabbitMQ) ConsumeSbuscription() {
	//第一步，试探性创建交换机exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建交换机失败。")

	//第二步，试探性创建队列queue
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建创建队列失败。")

	//第三步，绑定队列到交换机中
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下key要为空
		r.Exchange,
		false,
		nil,
	)

	//第四步，消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("订阅模式收到的消息：%s\n", d.Body)
		}
	}()

	log.Println("订阅模式消费者已开启，退出请按 CTRL+C")
	<-forever

}

//rabbitmq的路由模式。
//主要特点不仅一个消息可以被多个消费者消费还可以由生产端指定消费者。
//这里相对比订阅模式就多了一个routingkey的设计，也是通过这个来指定消费者的。
//创建exchange的kind需要是"direct",不然就不是roting模式了。

//创建rabbitmq实例，这里有了routingkey为参数了。
func NewRabbitMqRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建rabbit的路由实例的时候连接出现问题")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的路由实例时获取channel出错")
	return rabbitmq
}

//路由模式，产生消息。
func (r *RabbitMQ) PublishRouting(message string) {
	//第一步，尝试创建交换机，与pub/sub模式不同的是这里的kind需要是direct
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "路由模式，尝试创建交换机失败")
	//第二步，发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//路由模式，消费消息。
func (r *RabbitMQ) ConsumerRouting() {
	//第一步，尝试创建交换机，注意这里的交换机类型与发布订阅模式不同，这里的是direct
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建交换机失败。")

	//第二步，尝试创建队列,注意这里队列名称不用写，这样就会随机产生队列名称
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "路由模式，创建队列失败。")

	//第三步，绑定队列到exchange中
	err = r.channel.QueueBind(q.Name, r.Key, r.Exchange, false, nil)

	//第四步，消费消息。
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("路由模式(routing模式)收到消息为：%s。\n", d.Body)
		}
	}()
	<-forever
}

//topic模式
//与routing模式不同的是这个exchange的kind是"topic"类型的。
//topic模式的特别是可以以通配符的形式来指定与之匹配的消费者。
//"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。

//创建rabbitmq实例
func NewRabbitMqTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建rabbit的topic模式时候连接出现问题")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "创建rabbitmq的topic实例时获取channel出错")
	return rabbitmq
}

//topic模式。生产者。
func (r *RabbitMQ) PublishTopic(message string) {
	//第一步，尝试创建交换机,这里的kind的类型要改为topic
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式尝试创建exchange失败。")

	//第二步，发送消息。
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//topic模式。消费者。"*"表示匹配一个单词。“#”表示匹配多个单词，亦可以是0个。
func (r *RabbitMQ) ConsumerTopic() {
	//第一步，创建交换机。这里的kind需要是“topic”类型。
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true, //这里需要是true
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建exchange失败。")

	//第二步，创建队列。这里不用写队列名称。
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "topic模式，消费者创建queue失败。")

	//第三步，将队列绑定到交换机里。
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	//第四步，消费消息。
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			log.Printf("topic模式收到了消息：%s。\n", d.Body)
		}
	}()
	<-forever

}
