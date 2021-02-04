package main

import (
	"fmt"
	"log"
	"net"

	grpc_user "/gen/proto"
	server "/grpc"
	"/handler"

	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	port = ":50051"
)

func dbConnect(config) *gorm.DB {
	connectURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(config.DBType, connectURL)

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := dbConnect()
	handler := handler.NewHandler(db)

	s := grpc.NewServer()
	grpc_user.RegisterUserServiceServer(s, server.NewServer(handler))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
