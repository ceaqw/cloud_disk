package req

type LoginBody struct {
	Email    string `json:"email"`    //手机号
	Password string `json:"password"` //密码
}
