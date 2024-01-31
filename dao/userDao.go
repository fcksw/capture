package dao

import (
	"capture/initialize"
	"capture/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{initialize.NewDbClient(ctx)}
}


//一般应用于高级的transaction等操作 
func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{db}
}


func (dao *UserDao) QueryUserById(id int64) (result []*model.UserInfo, err error) {

	result = make([] *model.UserInfo, 0)
	dao.DB.Model(&model.UserInfo{}).Where("id > ?",0).Find(&result)
	return
}
