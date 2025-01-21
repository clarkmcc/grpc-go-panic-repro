package main

import (
	"context"
	_ "embed"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

const iterations = 1000

//go:embed image.png
var image []byte

func main() {
	if os.Args[1] == "grpc" {
		fmt.Println("running through grpc")
		runThroughGrpc()
	} else if os.Args[1] == "local" {
		fmt.Println("running locally")
		runLocally()
	}
}

func runThroughGrpc() {
	conn, err := grpc.NewClient("127.0.0.1:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	for i := 0; i < iterations; i++ {
		print(".")
		_, err = NewExampleClient(conn).Process(context.Background(), &ProcessRequest{
			Image: image,
		})
		if err != nil {
			panic(err)
		}
	}
}

func runLocally() {
	for i := 0; i < iterations; i++ {
		print(".")
		process(image)
	}
}
