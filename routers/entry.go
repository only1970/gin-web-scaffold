package routers

import (
	"fmt"
	"gvd_server/global"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	ginmode := global.Config.System.ENV

	if global.Config.System.ENV == "" {
		ginmode = "debug"
	}
	gin.SetMode(ginmode)

	router := gin.Default()
	fmt.Printf("%s 启动成功!！\n", global.Config.Logrus.AppName)
	return router
}
