# E2E-ServiceMesh-Test-Framework
## Go version
- go@1.15
- go@1.16
- go@1.17
- go@1.18

## STEPS
```zsh
# get into server folder of each rpc framework folder
go run ./xxx_server.go

# get into client folder of each rpc framework folder
go run ./xxx_client.go -c xxx -n xxx
```
```-s``` 指定服务端地址
```-c``` 指定并发数. 客户端会启动n个goroutine并发访问
```-n``` 指定测试的总请求数
