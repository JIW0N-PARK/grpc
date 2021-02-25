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
	VALID RegLogResult = 7;
)

type RegLogRes struct {
	Code RegLogResult
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Register(name string, email string, password string) (handlerRes *RegLogRes){
	user := model.FindByEmail(h.DB, email)
	valEmail := util.EmailCheck(email)
	valPassword := util.PasswordCheck(password)

	if user.Email == email {
		return &RegLogRes{Code: USER_REG_DUP_EMAIL}
	}
	if !valEmail {
		return &RegLogRes{Code: USER_REG_INVALID_EMAIL}
	}
	if !valPassword {
		return &RegLogRes{Code: USER_REG_INVALID_PASSWORD}
	}

	res := model.Insert(h.DB, name, email, password)

	if !res {
		return &RegLogRes{Code: USER_REG_FAIL}
	}

	return &RegLogRes{Code : USER_REG_SUCCESS}
}

func (h *UserHandler) Login(email string, password string) (handlerRes *RegLogRes) {
	user := model.FindByEmail(h.DB, email)

	if user.Email != email {
		return &RegLogRes{Code: USER_REG_INVALID_EMAIL}
	}

	if util.CompareHash(user.Password, password) != nil {
		return &RegLogRes{Code: USER_REG_INVALID_PASSWORD}
	}

	return &RegLogRes{Code: USER_LOG_SUCCESS}
}

func (h *UserHandler) ValidateEmail(email string) (handlerRes *RegLogRes) {
	user := model.FindByEmail(h.DB, email)
	valEmail := util.EmailCheck(email)

	if user.Email == email {
		return &RegLogRes{Code: USER_REG_DUP_EMAIL}
	}
	if !valEmail {
		return &RegLogRes{Code: USER_REG_INVALID_EMAIL}
	}
	return &RegLogRes{Code: VALID}
}

func (h *UserHandler) ValidatePassword(password string) (handlerRes *RegLogRes) {
	valPassword := util.PasswordCheck(password)

	if !valPassword {
		return &RegLogRes{Code: USER_REG_INVALID_PASSWORD}
	}
	return &RegLogRes{Code: VALID}
}
