package user

import (
	"net/http"
	common "project-common"
	"project-common/errs"
	"project-user/pkg/dao"
	"project-user/pkg/repo"
	loginServiceV1 "project-user/pkg/service/login.service.v1"
	"time"

	"context"

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
	result := &common.Result{}
	//1.获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
