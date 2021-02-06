package model

import (
	grpc_user "../../usersys/gen/proto"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	result := db.Where("name = ? AND age = ?", req.GetName(), req.GetAge()).Find(&User{})
	// db.Query("SELECT * FROM GRPC WHERE NAME = ?", name)

	name := req.GetName()
	age := req.GetAge()

	if result.RowsAffected == 1 {
		return fmt.Sprintf("%s(%d) is Found", name, age)
	} else {
		return fmt.Sprintf("%s(%d) is Not Found", name, age)

	}
}
