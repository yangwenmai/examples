package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/yangwenmai/examples/grpc/pb"
)

type routeGuideServer struct {
	savedFeatures []*pb.Feature
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

func (s *routeGuideServer) loadFeatures(jsonFile string) {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("failed to load default features:%v", err)
	}
	if err := json.Unmarshal(file, &s.savedFeatures); err != nil {
		log.Fatalf("failed to load default features:%v", err)
	}
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	s.loadFeatures(*jsonFile)
	return s
}

var (
	jsonFile = flag.String("jsonFile", "testdata/route_guide_db.json", "")
)

func main() {
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
