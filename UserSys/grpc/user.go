package user;

import (
	"context"

	grpc_user "../gen/proto"
	"../handler"

)

type Server struct {
	handler *handler.UserHandler
}

func NewServer(handler *handler.UserHandler) (s *Server) {
	return (
		&Server{handler: handler}
	)
}

func (s *Server) Register(ctx context.Context, req *grpc_user.Request) (res *grpc_user.Response, error) {
	res := s.handler.Register(req)
	return &grpc_user.Response{res: res}, nil
}

func (s *Server) Search(ctx context.Context, req *grpc_user.Request) (res *grpc_user.Response, error) {
	res := s.handler.Search(req)
	return &grpc_user.Response{res: res}, nil
}