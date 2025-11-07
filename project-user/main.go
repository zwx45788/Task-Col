package main

import (
	commom "project-common"
	_ "project-user/api"
	"project-user/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	commom.Run(r, "project-user", ":80")

}
