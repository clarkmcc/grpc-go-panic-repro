package main

import (
	"context"
	_ "embed"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:embed image.png
var image []byte

func main() {
	conn, err := grpc.NewClient("127.0.0.1:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		print(".")
		_, err = NewExampleClient(conn).Process(context.Background(), &ProcessRequest{
			Image: image,
		})
		if err != nil {
			panic(err)
		}
	}
}
