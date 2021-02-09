package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int32
	Name string
	Email  string
	Password string
}

func Insert(db *gorm.DB, name string, email string, password string) bool {
	user := User{Name: name, Email: email, Password: password}

	result := db.Create(&user)

	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func Find(db *gorm.DB, email string) bool {
	result := db.Where("email = ?", email).Find(&User{})

	if result.RowsAffected == 1 {
		//find
		return true
	} else {
		return false
	}
}

func Match(db *gorm.DB, email string, password string) bool {
	result := db.Where("email = ? and password = ?", email, password).Find(&User{})

	if result.RowsAffected == 1 {
		//login
		return true
	} else {
		return false
	}
}

