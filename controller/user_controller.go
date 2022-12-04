package controller

import (
	"CouldDisk/middleware/jwt"
	"CouldDisk/models/req"
	"CouldDisk/models/resp"
	"CouldDisk/services"
	"CouldDisk/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userService services.UserService
}

func (u User) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		result := make(map[string]interface{})
		login, token, user := u.userService.Login(loginBody.Email, loginBody.Password)
		if login {
			result["token"] = token
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

func (u User) CheckUserLoginInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, resp.CheckTokenError())
		return
	}
	j := jwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, resp.CheckTokenError())
		return
	}
	data := make(map[string]interface{})
	data["username"] = claims.UserInfo.Name
	data["userid"] = claims.UserInfo.Id
	c.JSON(http.StatusOK, resp.CheckTokenSuccess(data))
}

func (u User) Register(c *gin.Context) {
	registerBody := req.RegisterBody{}
	if c.BindJSON(&registerBody) == nil {
		fmt.Println(registerBody)
		isRegister, msg := u.userService.Register(registerBody.Name, registerBody.Email, registerBody.Password)
		c.JSON(200, resp.BoolResponse(isRegister, msg))
	} else {
		resp.ParamError(c)
	}
}

func (u User) UpdatePwdByEmail(c *gin.Context) {
	email := c.Query("email")
	verifyCode := c.Query("verifyCode")
	newPwd := c.Query("newPwd")
	b := u.userService.UpdatePwdByEmail(email, newPwd, verifyCode)
	if b {
		c.JSON(200, resp.Success(nil))
	} else {
		c.JSON(200, resp.ErrorResp())
	}
}

func (u User) SendCode(c *gin.Context) {
	email := c.Query("email")
	err := utils.SendVerifyCode(email)
	if err != nil {
		c.JSON(200, resp.ErrorResp(err))
	} else {
		c.JSON(200, resp.Success(nil))
	}
}
