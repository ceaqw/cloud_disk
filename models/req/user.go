package req

type LoginBody struct {
	Email    string `json:"email"`    //邮箱
	Password string `json:"password"` //密码
}

type RegisterBody struct {
	Name     string `json:"name"`     //姓名
	Email    string `json:"email"`    //邮箱
	Password string `json:"password"` //密码
}
