package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int32 `gorm:"primaryKey"`
	Name string
	Age  int32
}

func Insert(db *gorm.DB, req *grpc_user.Request) string {
	var count int32
	db.Model(&User{}).Count(&count)
	user := User{Id: count + 1, Name: req.GetName(), Age: req.GetAge()}

	result := db.Create(&user)

	if result.RowsAffected == 1 {
		return "Complete Insert!"
	} else {
		return "Failed.."
	}
}

func Find(db *gorm.DB, req *grpc_user.Request) string {
	result := db.Where("name = ? AND age = ?", req.GetName(), req.GetAge()).Find(&User)
	// db.Query("SELECT * FROM GRPC WHERE NAME = ?", name)

	if result.RowsAffected == 1 {
		return name + " is Found"
	} else {
		return name + " is Not Found"
	}
}
