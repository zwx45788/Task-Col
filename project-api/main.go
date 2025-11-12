package main

import (
	common "project-common"
	_ "project-user/api"
	"project-user/config"

	"project-user/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	// gc := router.RegisterGrpc()
	// stop := func() {
	// 	gc.Stop()
	// }
	common.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)

}
