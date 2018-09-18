package main

import (
	"context"
	"log"

	"github.com/yangwenmai/examples/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dail :%v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	feature, err := client.GetFeature(context.Background(), &pb.Point{Latitude: 409146138, Longitude: -746188906})
	if err != nil {
		log.Fatalf("client failed:%v", err)
	}
	log.Println(feature)
}
