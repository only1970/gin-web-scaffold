package flags

import "gin-web-scaffold/global"

func Port(port int) {
	if port == 0 {
		global.Log.Fatal("listen tcp: address 0: invalid port")
	}
	global.Config.System.Port = port
}
