import (
	"../model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) (*UserHandler){
	return (
		&UserHandler{DB: db}
	)
}

func (h *UserHandler) Register(req *grpc_user.Request) string {
	res := model.Insert(h.DB, req)
	return res
}

func (h *UserHandler) Search(req *grpc_user.Request) string {
	res := model.Find(h.DB, req)
	return res
}
