package user

import (
	"net/http"
	common "project-common"
	"project-common/errs"

	login "project-grpc/user/login"
	"time"

	"context"

	"project-api/pkg/model/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type HandlerUser struct {
}

func New() *HandlerUser {

	return &HandlerUser{}
}
func (hl *HandlerUser) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	//1.获取参数
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp, err := LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
func (u *HandlerUser) Register(c *gin.Context) {
	result := &common.Result{}
	//1.接收参数 参数模型
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式错误"))
		return
	}
	//2.校验参数 判断参数是否合法
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	//3.调用user grpc服务 获取响应
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "copy错误"))
		return
	}
	_, err = LoginServiceClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	//4.返回响应
	c.JSON(http.StatusOK, result.Success("注册成功"))

}
