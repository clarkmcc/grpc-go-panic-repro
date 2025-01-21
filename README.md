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
Serial Number (system): FRX9W6F6V1
Hardware UUID: 0BBEE97F-CEA2-577D-B308-DBF11EB318D5
Provisioning UDID: 00006001-001C21A11A02801E
Activation Lock Status: Enabled
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
Serial Number (system): VKL7N67RWF
Hardware UUID: 258890EE-B06B-5109-81CC-6CAA34E65F12
Provisioning UDID: 00008132-000A109A3EC3001C
Activation Lock Status: Disabled
```