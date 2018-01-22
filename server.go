package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/alextanhongpin/go-twirp/proto"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}

func main() {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)

	mux := http.NewServeMux()

	mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)

	log.Println("Listening to port *:8080. Press ctrl + c to cancel.")
	http.ListenAndServe(":8080", mux)
}
