package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "github.com/yangwenmai/examples/tts-grpc-k8s/pb"
	grpc "google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	log.Printf("listening to port: %v\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("listen failed :%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("could not to serve:%v", err)
	}
}

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.Text) (*pb.Speech, error) {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("could not close :%v", err)
	}

	log.Printf("file name:%v", file.Name())
	cmd := exec.Command("flite", "-t", in.Text, "-o", file.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed %s", data)
	}

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read temp file:%v", err)
	}
	return &pb.Speech{Audio: data}, nil
}

// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }
