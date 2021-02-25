package server

import (
	"context"
	grpc "grpc_ex/gen/proto"
	"grpc_ex/handler"
)

type Server struct {
	grpc.UnimplementedUserServiceServer
	handler *handler.UserHandler
}

func NewServer(handler *handler.UserHandler) *Server {
	return &Server{handler: handler}
}

var (
	ResultMap = map[handler.RegLogResult]grpc.ResponseResult{
		handler.USER_REG_SUCCESS: grpc.ResponseResult_REGISTERED,
		handler.USER_REG_DUP_EMAIL: grpc.ResponseResult_DUPLICATED_EMAIL,
		handler.USER_REG_INVALID_EMAIL: grpc.ResponseResult_INVALID_EMAIL,
		handler.USER_REG_INVALID_PASSWORD: grpc.ResponseResult_INVALID_PASSWORD,
		handler.USER_REG_FAIL: grpc.ResponseResult_NOT_REGISTERED,
		handler.USER_LOG_SUCCESS: grpc.ResponseResult_LOGIN_SUCCESS,
		handler.VALID: grpc.ResponseResult_VALID,
	}
)

func (s *Server) Register(ctx context.Context, req *grpc.RegisterReq) (*grpc.Res, error) {
	handlerRes := s.handler.Register(req.Name, req.Email, req.Password)

	return &grpc.Res{Response: ResultMap[handlerRes.Code]}, nil
}

func (s *Server) Login(ctx context.Context, req *grpc.LoginReq) (*grpc.Res, error) {
	handlerRes := s.handler.Login(req.Email, req.Password)

	return &grpc.Res{Response: ResultMap[handlerRes.Code]}, nil
}

func (s *Server) ValidateEmail(ctx context.Context, req *grpc.EmailReq) (*grpc.Res, error){
	handlerRes := s.handler.ValidateEmail(req.Email)

	return &grpc.Res{Response: ResultMap[handlerRes.Code]}, nil
}

func (s *Server) ValidatePassword(ctx context.Context, req *grpc.PasswordReq) (*grpc.Res, error){
	handlerRes := s.handler.ValidatePassword(req.Password)

	return &grpc.Res{Response: ResultMap[handlerRes.Code]}, nil
}
