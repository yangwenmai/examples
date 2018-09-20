package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yangwenmai/examples/grpc-k8s-example/pb"
	"google.golang.org/grpc"
)

func main() {
	a := flag.Uint64("a", 3, "a")
	b := flag.Uint64("b", 4, "b")
	flag.Parse()

	conn, err := grpc.Dial("192.168.99.100:30265", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	gcdClient := pb.NewGCDServiceClient(conn)

	// Call GCD service
	req := &pb.GCDRequest{A: *a, B: *b}
	if res, err := gcdClient.Compute(context.Background(), req); err == nil {
		fmt.Printf("result:%d, version:%s", res.Result, res.Version)
	} else {
		fmt.Println("error...")
	}
}
