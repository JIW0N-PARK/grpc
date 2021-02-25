package model

import (
	"github.com/jinzhu/gorm"
	"grpc_ex/util"
)

type User struct {
	Id   int32
	Name string
	Email  string
	Password string
}

func Insert(db *gorm.DB, name string, email string, password string) bool {
	pw, err := util.GenerateHash(password)

	if err != nil {
		return false
	}

	user := User{Name: name, Email: email, Password: string(pw)}

	result := db.Create(&user)

	return result.RowsAffected == 1
}

func FindByEmail(db *gorm.DB, email string) *User {
	user := &User{}
	db.Where("email = ?", email).Find(user)

	return user
}


