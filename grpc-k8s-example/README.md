# gRPC-k8s-example

## pb

protocol buffer 文件

## gcd

gcd 服务

## api

提供 API 的方式来调用。

*问题：*

- API 内部调用 gcd 服务，无法使用到 gcd 的 service 负载均衡。
- 使用 kuberesolver 后 gcd 服务依然不能做到负载均衡。

## client

本地直接执行 `go run client.go`，这样运行因为使用到的是 k8s loadbalancer ，所以负载均衡能够生效。

## Drone

本项目是用于 Gitlab + Drone + k8s 构建 CI/CD 的演示项目。

**注意：项目受限于环境配置的问题，所以需要自己调整一下目录和配置对应参数才能运行。有时间后将其修复。**

该项目来源于：[Getting Started with Microservices using Go, gRPC and Kubernetes](https://outcrawl.com/getting-started-microservices-go-grpc-kubernetes/)
