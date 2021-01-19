package main

import (
	"context"
	"net"
	"log"

	"google.golang.org/grpc"
	pb "../proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRes, error){
	log.Printf("received rpc from client, name=%s\n", in.GetReq())
	return &pb.HelloRes{Res: "Hello" + in.GetReq()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}