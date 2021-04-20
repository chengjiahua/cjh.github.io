package main

import (
	"flag"
	"fmt"
	RabbitMq "rabbitmq-example/RabbitMQ"
	"strconv"
	"time"
)

var mode = flag.Int64("mode", 0, "0为简单模式，1为订阅模式，2为路由模式，3为主题模式")

func main() {
	flag.Parse()
	switch *mode {
	case 0:
		Simple()
	case 1:
		Subscribe()
	case 2:
		Routing()
	case 3:
		Topic()
	default:
		fmt.Println("This mode does not exist")
	}
}

func Simple() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("SimpleMode")
	rabbitmq.PublishSimple("Hello World ---来自simple模式")
	fmt.Println("发送成功！")

}

func Subscribe() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("SubscribeMode")
	for i := 0; i < 20; i++ {
		rabbitmq.PublishSubscription("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
		time.Sleep(1 * time.Second)
	}
}

func Routing() {
	rabbitmq1 := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "one")
	rabbitmq2 := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "two")
	rabbitmq3 := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "three")
	for i := 0; i < 100; i++ {
		rabbitmq1.PublishRouting("路由模式one" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("路由模式two" + strconv.Itoa(i))
		rabbitmq3.PublishRouting("路由模式three" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("在路由模式下，routingKey为one,为two,为three的都分别生产了%d条消息\n", i)
	}

}

func Topic() {
	one := RabbitMq.NewRabbitMqTopic("TopicMode", "Singer.Jay")
	two := RabbitMq.NewRabbitMqTopic("TopicMode", "Persident.XIDADA")
	for i := 0; i < 100; i++ {
		one.PublishTopic("topic模式，Jay," + strconv.Itoa(i))
		two.PublishTopic("topic模式，All," + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Printf("topic模式发布的消息%v \n", i)
	}
}
