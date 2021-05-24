package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pushlish "grpc/pushlish"

	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pushlish.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &pushlish.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
