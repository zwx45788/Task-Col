package main

import (
	"log"
	common "project-common"
	"project-common/logs"
	_ "project-user/api"
	"project-user/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	lc := &logs.LogConfig{
		DebugFileName: "D:\\git-demo\\ms-project\\logs\\debug",
		InfoFileName:  "D:\\git-demo\\ms-project\\logs\\info",
		WarnFileName:  "logs/user_warn.log",
		MaxSize:       500, // MB
		MaxAge:        28,  // Days
		MaxBackups:    3,
	}

	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	common.Run(r, config.C.SC.Name, config.C.SC.Addr)

}
