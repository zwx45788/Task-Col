package user

import (
	"errors"
	common "project-common"
)

type RegisterReq struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

func (r RegisterReq) Verify() error {
	if !common.VerifyEmail(r.Email) {
		return errors.New("邮箱不合法")
	}
	if !common.VerifyPassword(r.Password, r.Password2) {
		return errors.New("密码不匹配")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号不合法")
	}
	return nil
}
