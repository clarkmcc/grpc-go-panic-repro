Reproduction for https://github.com/grpc/grpc-go/issues/8023.

## Steps to reproduce
Start the server
```shell
go run server.go example.pb.go example_grpc.pb.go
```

Start the client
```shell
go run client.go example.pb.go example_grpc.pb.go
```