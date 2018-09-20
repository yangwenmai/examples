package main

import (
	"log"
	"net"

	"github.com/yangwenmai/examples/grpc-k8s-example/pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var VERSION = "0.1"

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	log.Printf("[%s], arguments a:%v, b:%v, Result:%v", VERSION, r.A, r.B, a)
	return &pb.GCDResponse{Result: a, Version: VERSION}, nil
}
