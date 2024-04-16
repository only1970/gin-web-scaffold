package core

import (
	"fmt"
	"gin-web-scaffold/config"
	"gin-web-scaffold/global"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "config.yml"

func InitConfig() (c *config.Config) {
	bytedata, err := os.ReadFile(configPath)
	if err != nil {
		global.Log.Fatalln("read yaml err: ", err.Error())
	}
	c = new(config.Config)
	err = yaml.Unmarshal(bytedata, c)
	if err != nil {
		global.Log.Fatalln("yaml unmarshal err: ", err.Error())
	}
	fmt.Printf("配置文件初始化成功 %v \n", *c)
	return c

}
