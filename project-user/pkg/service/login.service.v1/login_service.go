package login_service_v1

import (
	"context"
	"log"
	common "project-common"
	"project-common/errs"
	login "project-grpc/user/login"
	"project-user/internal/dao"
	"project-user/internal/repo"
	"project-user/pkg/model"
	"time"

	"go.uber.org/zap"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	Cache      repo.Cache
	memberRepo repo.MemberRepo
}

func NewLoginService() *LoginService {
	return &LoginService{
		Cache:      dao.Rc,
		memberRepo: dao.NewMemberDao(),
	}
}
func (ls *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {

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
		err := ls.Cache.Put(model.RedisKeyRegisterCaptcha+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码发生错误,cause by:", err)
		}
		//redis.Set("REGISTER_"+mobile, code)
		log.Printf("将手机号和验证码存入redis成功:REGISTER_%s : %s", mobile, code)
	}()

	//ctx.JSON(http.StatusOK, rsp.Success("123456"))
	return &login.CaptchaResponse{}, nil
}
func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	c := context.Background()
	//1.校验参数
	//2.校验验证码
	redisCode, err := ls.Cache.Get(model.RedisKeyRegisterCaptcha + msg.Mobile)
	if err != nil {
		zap.L().Error("获取验证码失败", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != msg.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	//3.校验业务逻辑
	exist, err := ls.memberRepo.GetMemberByEmail(c, msg.Email)
	if err != nil {
		zap.L().Error("数据库错误", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	exist, err = ls.memberRepo.GetMemberByMobile(c, msg.Mobile)
	if err != nil {
		zap.L().Error("数据库错误", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.MobileExist)
	}
	//4.入库
	//5.返回响应
	return &login.RegisterResponse{}, nil
}
