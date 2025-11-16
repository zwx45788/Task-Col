package user

import (
	"log"
	"net/http"
	common "project-common"
	"project-user/internal/dao"
	"project-user/internal/repo"
	"project-user/pkg/model"
	"time"

	"project-common/errs"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {

	return &HandlerUser{
		cache: dao.Rc,
	}
}
func (hl *HandlerUser) GetCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	//1.获取参数
	mobile := ctx.PostForm("mobile")
	//2.校验参数
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, errs.GrpcError(model.NoLegalMobile))
		return
	}
	//3.生成验证码
	code := "123456"
	//4.调用短信平台
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信平台调用成功,发送短信")
		//5.存储验证码redis到中,过期时间15min
		err := hl.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码发生错误,cause by:", err)
		}
		//redis.Set("REGISTER_"+mobile, code)
		log.Printf("将手机号和验证码存入redis成功:REGISTER_%s : %s", mobile, code)
	}()

	ctx.JSON(http.StatusOK, rsp.Success("123456"))
}
