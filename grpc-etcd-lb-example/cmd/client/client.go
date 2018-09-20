package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"gogs.maiyang.me/developer-learning/grpc-etcd-lb/cmd/pb"
	grpclb "gogs.maiyang.me/developer-learning/grpc-etcd-lb/etcdv3"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	service    = flag.String("service", "hello_service", "service name")
	etcdServer = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()
	r := grpclb.NewResolver(*service)
	b := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *etcdServer, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world" + strconv.Itoa(t.Second())})
		if err != nil {
			fmt.Printf("%v: Reply is %s\n", resp.Message)
		}
		fmt.Printf("%v : Reply is %v\n", time.Now(), resp.Message)
	}
}
