package service

import (
	"capture/dao"
	"capture/model"
	"context"
	"sync"
)


var UserServiceIns *UserService
var UserSrvOnce sync.Once

type UserService struct {

}

func GetUserServiceIns() *UserService{
	UserSrvOnce.Do(func() {
		UserServiceIns = &UserService{}
	})
	return UserServiceIns
}


func (service *UserService) ListUser(ctx context.Context) (result []*model.UserInfo, err error){
	return dao.NewUserDao(ctx).QueryUserById(0);
}








