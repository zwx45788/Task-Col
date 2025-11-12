package login_service_v1

import (
	"context"
	"log"
	common "project-common"
	"project-common/errs"
	"project-user/pkg/dao"
	"project-user/pkg/model"
	"project-user/pkg/repo"

	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	Cache repo.Cache
}

func NewLoginService() *LoginService {
	return &LoginService{
		Cache: dao.Rc,
	}
}
func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {

	//1.获取参数
	mobile := msg.Mobile
	//2.校验参数
	if !common.VerifyMobile(mobile) {
		//ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobile, "手机号不合法"))
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	//3.生成验证码
	code := "123456"
	//4.调用短信平台
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信平台调用成功,发送短信")
		//5.存储验证码redis到中,过期时间15min
		err := ls.Cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码发生错误,cause by:", err)
		}
		//redis.Set("REGISTER_"+mobile, code)
		log.Printf("将手机号和验证码存入redis成功:REGISTER_%s : %s", mobile, code)
	}()

	//ctx.JSON(http.StatusOK, rsp.Success("123456"))
	return &CaptchaResponse{}, nil
}
