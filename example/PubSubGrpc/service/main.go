package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	pushlish "grpc/pushlish"

	"github.com/moby/moby/pkg/pubsub"
	grpc "google.golang.org/grpc"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, arg *pushlish.String,
) (*pushlish.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pushlish.String{}, nil
}

func (p *PubsubService) Subscribe(
	arg *pushlish.String, stream pushlish.PubsubService_SubscribeServer,
) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pushlish.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	p := grpc.NewServer()

	pushlish.RegisterPubsubServiceServer(p, NewPubsubService())
	p.Serve(listener)
}
