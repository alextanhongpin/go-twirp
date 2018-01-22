package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/alextanhongpin/go-twirp/proto"
)

func main() {
	client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
	if err != nil {
		log.Println(resp.Text)
	}
	log.Println(resp.Text)
}
