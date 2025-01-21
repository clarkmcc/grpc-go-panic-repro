package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

type impl struct {
	UnimplementedExampleServer
}

func main() {
	fmt.Println(os.Getpid())
	srv := grpc.NewServer()
	var i impl
	RegisterExampleServer(srv, &i)
	l, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(l)
	if err != nil {
		panic(err)
	}
}

func (i *impl) Process(_ context.Context, req *ProcessRequest) (*ProcessResponse, error) {
	process(req.Image)
	return &ProcessResponse{}, nil
}
