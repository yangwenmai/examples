package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/yangwenmai/examples/tts-grpc-k8s/api"
	grpc "google.golang.org/grpc"
)

func main() {
	server := flag.String("b", "192.168.99.100:30245", "address of say backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speech\"\n", os.Args[0])
		os.Exit(1)
	}

	con, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial err:%v", err)
	}
	defer con.Close()

	client := pb.NewTextToSpeechClient(con)
	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could to say %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("write file err:%v", err)
	}
}
