package main

import (
	"gvd_server/core"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	// global.Log = core.InitLogger(core.LogRequest{
	// 	LogPath: "./newlogs",
	// 	AppName: "gvd_server",
	// })
	global.Config = core.InitConfig()
	global.Log = core.InitLogger(&global.Config.Logrus)
	global.MySQL = core.InitGorm()
	global.Redis = core.InitRedis()

	option := flags.Parse()
	if option.RunOptions() {
		return
	}

	routers := routers.Routers()
	routers.Run(global.Config.System.Addr())
}
