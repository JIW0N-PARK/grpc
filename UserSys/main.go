package main

import (
	"log"
	"net"

	grpc_user "../usersys/gen/proto"
	datab "../usersys/database"
	server "../usersys/grpc"
	"../usersys/handler"

	"google.golang.org/grpc"
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

	db := datab.DBInit(mysql)
	newHandler := handler.NewHandler(db)

	s := grpc.NewServer()
	grpc_user.RegisterUserServiceServer(s, server.NewServer(newHandler))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
