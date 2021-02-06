package server

import (
	grpc_user "../../usersys/gen/proto"
	"../../usersys/handler"
	"context"
	"log"
)

type Server struct {
	grpc_user.UnimplementedUserServiceServer
	handler *handler.UserHandler
}

func NewServer(handler *handler.UserHandler) *Server {
	return &Server{handler: handler}
}

func (s *Server) Register(ctx context.Context, req *grpc_user.Request) (*grpc_user.Response, error) {
	res := s.handler.Register(req)
	log.Printf("Register : %s", res.Res)
	return res, nil
}

func (s *Server) Search(ctx context.Context, req *grpc_user.Request) (*grpc_user.Response, error) {
	res := s.handler.Search(req)
	log.Printf("Search : %s", res.Res)
	return res, nil
}
