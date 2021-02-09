package server

import (
	"context"
	grpc_user "grpc_ex/gen/proto"
	"grpc_ex/handler"
)

type Server struct {
	grpc_user.UnimplementedUserServiceServer
	handler *handler.UserHandler
}

func NewServer(handler *handler.UserHandler) *Server {
	return &Server{handler: handler}
}

var (
	ResultMap = map[handler.RegLogResult]grpc_user.ResponseResult{
		handler.USER_REG_SUCCESS: grpc_user.ResponseResult_REGISTERED,
		handler.USER_REG_DUP_EMAIL: grpc_user.ResponseResult_DUPLICATED_EMAIL,
		handler.USER_REG_INVALID_EMAIL: grpc_user.ResponseResult_INVALID_EMAIL,
		handler.USER_REG_INVALID_PASSWORD: grpc_user.ResponseResult_INVALID_PASSWORD,
		handler.USER_REG_FAIL: grpc_user.ResponseResult_NOT_REGISTERED,
		handler.USER_LOG_SUCCESS: grpc_user.ResponseResult_LOGIN_SUCCESS,
	}
)

func (s *Server) Register(ctx context.Context, req *grpc_user.Request) (*grpc_user.Response, error) {
	handlerRes := s.handler.Register(req.Name, req.Email, req.Password)

	return &grpc_user.Response{Response: ResultMap[handlerRes.Code]}, nil
}

func (s *Server) Login(ctx context.Context, req *grpc_user.Request) (*grpc_user.Response, error) {
	handlerRes := s.handler.Login(req.Email, req.Password)

	return &grpc_user.Response{Response: ResultMap[handlerRes.Code]}, nil
}
