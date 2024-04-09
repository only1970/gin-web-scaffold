package flags

import (
	"flag"

	"github.com/bytedance/sonic/option"
)

type Option struct {
	DB   bool   //初始化数据库
	Port int    //指定端口
	Load string //导入数据库文件

}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.IntVar(&option.Port, "port", 8080, "指定端口")
	flag.StringVar(&option.Load, "load", "", "导入sql文件")
	flag.Parse()
	return option
}

//根据不同的参数运行不同的脚本

func (option Option)RunOptions() bool  {
	return false
}