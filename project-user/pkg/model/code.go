package model

import (
	"project-common/errs"
)

var (
	NoLegalMobile = errs.NewError(10102001, "手机号不合法")
	CaptchaError  = errs.NewError(10102002, "验证码错误")
	EmailExist    = errs.NewError(10102003, "邮箱已存在")
	MobileExist   = errs.NewError(10102004, "手机号已存在")
	RedisError    = errs.NewError(-100, "Redis操作失败")
	DBError       = errs.NewError(-200, "DB错误")
)
