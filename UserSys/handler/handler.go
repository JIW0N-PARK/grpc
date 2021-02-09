package handler

import (
	"github.com/jinzhu/gorm"
	"grpc_ex/model"
	"grpc_ex/util"
)

type UserHandler struct {
	DB *gorm.DB
}

type RegLogResult int

const (
	USER_REG_SUCCESS RegLogResult = 1;
	USER_REG_DUP_EMAIL RegLogResult = 2;
	USER_REG_INVALID_EMAIL RegLogResult = 3;
	USER_REG_INVALID_PASSWORD RegLogResult = 4;
	USER_REG_FAIL RegLogResult = 5;
	USER_LOG_SUCCESS RegLogResult = 6;
)

type RegLogRes struct {
	Code RegLogResult
}

func NewHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Register(name string, email string, password string) (handlerRes *RegLogRes){
	dupEmail := model.Find(h.DB, email)
	valEmail := util.EmailCheck(email)
	valPassword := util.PasswordCheck(password)

	if dupEmail {
		return &RegLogRes{Code: USER_REG_DUP_EMAIL}
	}
	if ! valEmail {
		return &RegLogRes{Code: USER_REG_INVALID_EMAIL}
	}
	if ! valPassword {
		return &RegLogRes{Code: USER_REG_INVALID_PASSWORD}
	}

	res := model.Insert(h.DB, name, email, password)

	if res {
		return &RegLogRes{Code : USER_REG_SUCCESS}
	}else {
		return &RegLogRes{Code: USER_REG_FAIL}
	}


}

func (h *UserHandler) Login(email string, password string) (handlerRes *RegLogRes) {
	emailCheck := model.Find(h.DB, email)

	if ! emailCheck {
		return &RegLogRes{Code: USER_REG_INVALID_EMAIL}
	}

	login := model.Match(h.DB, email, password)

	if login {
		return &RegLogRes{Code: USER_LOG_SUCCESS}
	} else {
		return &RegLogRes{Code: USER_REG_INVALID_PASSWORD}
	}
}
