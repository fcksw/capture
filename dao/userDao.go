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


func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{db}
}



// func (dao *UserDao) createUser() (err error) {
// 	dao.DB.
// }


func (dao *UserDao) queryUserById(id int64) (result []*model.User, err error) {
	result = make([] *model.User, 0)
	dao.DB.Model(&model.User{}).Where("id > ",0).Find(&result)
	return
}
