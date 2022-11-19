package controller

import (
	"CouldDisk/models/req"
	"CouldDisk/models/resp"
	"CouldDisk/services"

	"github.com/gin-gonic/gin"
)

type User struct {
	userService services.UserService
}

func (u User) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(loginBody) != nil {
		result := make(map[string]interface{})
		login, token, user := u.userService.Login(loginBody.Email, loginBody.Password)
		if login {
			result["toekn"] = token
			result["userid"] = user.Id
			result["name"] = user.Name
			c.JSON(200, resp.Success(result))
		} else {
			c.JSON(200, resp.ErrorResp(token))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}
