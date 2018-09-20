package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gogs.maiyang.me/developer-learning/grpc-etcd-lb/cmd/pb"
	grpclb "gogs.maiyang.me/developer-learning/grpc-etcd-lb/etcdv3"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	service    = flag.String("service", "hello_service", "service name")
	host       = flag.String("host", "localhost", "listening host")
	port       = flag.String("port", "50001", "listening port")
	etcdServer = flag.String("etcdServer", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		panic(err)
	}

	err = grpclb.Register(*service, *host, *port, *etcdServer, 10*time.Second, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		grpclb.UnRegister()
		os.Exit(1)
	}()

	log.Printf("starting hello service at %s", *port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(listen)
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name + " from " + net.JoinHostPort(*host, *port)}, nil
}
