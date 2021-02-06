package handler

import (
	grpc_user "../../usersys/gen/proto"
	"../../usersys/model"

	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Register(req *grpc_user.Request) (*grpc_user.Response) {
	res := model.Insert(h.DB, req)
	return &grpc_user.Response{Res: res}
}

func (h *UserHandler) Search(req *grpc_user.Request) (*grpc_user.Response) {
	res := model.Find(h.DB, req)
	return &grpc_user.Response{Res: res}
}
