package main

import (
	"log"
	"net"

	grpc_user "github.com/jiohning/usersys/gen/proto"
	server "github.com/jiohning/usersys/grpc"
	"github.com/jiohning/usersys/handler"

	// "github.com/jiohning/usersys/database"

	"google.golang.org/grpc"
)

const (
	port     = ":50051"
	mysql    = "mysql"
	postgres = "postgres"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := database.dbConnect(mysql)
	handler := handler.NewHandler(db)

	s := grpc.NewServer()
	grpc_user.RegisterUserServiceServer(s, server.NewServer(handler))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
