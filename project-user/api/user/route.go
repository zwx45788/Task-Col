package user

import (
	"log"
	"project-user/router"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/project/login/getCaptcha", h.GetCaptcha)
}
