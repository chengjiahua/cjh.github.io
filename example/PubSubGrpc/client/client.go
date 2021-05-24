package main

import (
	"context"
	"log"

	pushlish "grpc/pushlish"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pushlish.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &pushlish.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &pushlish.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
