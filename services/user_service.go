package services

import (
	"CouldDisk/middleware/jwt"
	"CouldDisk/models"

	"github.com/druidcaesa/gotool"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userModel models.UserOrm
}

func (s UserService) Login(email string, password string) (bool, string, *models.UserBasic) {
	user := s.userModel.GetUserByEmail(email)
	// fmt.Println(user)
	if user == nil || !gotool.BcryptUtils.CompareHash(user.Password, password) {
		return false, "用户名或密码错误", user
	}
	//生成token
	token, err := jwt.NewJWT().CreateUserToken(*user)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return false, "", user
	}
	//数据存储到redis中
	return true, token, user
}
