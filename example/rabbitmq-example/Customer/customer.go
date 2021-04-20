package main

import (
	"flag"
	"fmt"
	RabbitMq "rabbitmq-example/RabbitMQ"
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
	rabbitmq.ConsumeSimple()
}

func Subscribe() {
	rabbitmq := RabbitMq.NewRabbitMqSubscription("SubscribeMode")
	rabbitmq.ConsumeSbuscription()
}

func Routing() {
	// customer1
	// one := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "one")
	// one.ConsumerRouting()

	// customer2
	// two := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "two")
	// two.ConsumerRouting()

	// customer3
	four := RabbitMq.NewRabbitMqRouting("RouteKeyMode", "four")
	four.ConsumerRouting()
}

func Topic() {

	// customer1
	jay := RabbitMq.NewRabbitMqTopic("TopicMode", "Singer.*")
	jay.ConsumerTopic()

	// customer2
	// jay := RabbitMq.NewRabbitMqTopic("TopicMode", "#")
	// jay.ConsumerTopic()
}
