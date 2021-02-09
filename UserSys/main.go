package main

import (
	"google.golang.org/grpc"
	"grpc_ex/database"
	grpc_user "grpc_ex/gen/proto"
	server "grpc_ex/grpc"
	"grpc_ex/handler"

	"log"
	"net"
)

const (
	port     = ":8080"
	mysql    = "mysql"
	postgres = "postgres"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := database.DBInit(mysql)
	newHandler := handler.NewHandler(db)

	s := grpc.NewServer()
	grpc_user.RegisterUserServiceServer(s, server.NewServer(newHandler))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
