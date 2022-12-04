package services

import (
	"CouldDisk/middleware/jwt"
	"CouldDisk/models"
	"CouldDisk/utils"

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

func (s UserService) Register(name, email, password string) (bool, string) {
	if s.userModel.GetUserByEmail(email) != nil {
		return false, "该邮箱已经被注册"
	}
	user_basic := models.UserBasic{}
	user_basic.Password = string(gotool.BcryptUtils.Generate(password))
	user_basic.Email = email
	user_basic.Name = name
	err := s.userModel.AddUser(user_basic)
	if err != nil {
		return false, "注册失败"
	}
	return true, "注册成功"
}

func (s UserService) UpdatePwdByEmail(email, password, verifyCode string) bool {
	ok := utils.CheckVerifyCode(email, verifyCode)
	if ok {
		crypassword := string(gotool.BcryptUtils.Generate(password))
		b := s.userModel.UpdatePwdByEmail(email, crypassword)
		return b
	}
	return false
}
