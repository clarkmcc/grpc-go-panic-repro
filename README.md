Reproduction for https://github.com/grpc/grpc-go/issues/8023.

## Steps to reproduce
Start the server
```shell
go run server.go example.pb.go example_grpc.pb.go vision.go
```

Start the client
```shell
go run client.go example.pb.go example_grpc.pb.go vision.go grpc
```

You can run the client without gRPC which will run the Apple Vision Framework calls locally rather than through gRPC. The crash is only reproduced over gRPC, the same Apple Vision Framework code works when executed directly from `client.go` rather than through the gRPC service.

```shell
go run client.go example.pb.go example_grpc.pb.go vision.go local
```

## Environments

Reproducible on the following machines
```
Model Name: MacBook Pro
Model Identifier: MacBookPro18,2
Model Number: Z14W0010BLL/A
Chip: Apple M1 Max
Total Number of Cores: 10 (8 performance and 2 efficiency)
Memory: 64 GB
System Firmware Version: 11881.41.5
OS Loader Version: 11881.41.5
```

```
Model Name: Mac mini
Model Identifier: Mac16,10
Model Number: MU9D3LL/A
Chip: Apple M4
Total Number of Cores: 10 (4 performance and 6 efficiency)
Memory: 16 GB
System Firmware Version: 11881.61.3
OS Loader Version: 11881.61.3
```