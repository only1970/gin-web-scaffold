package core

import (
	"fmt"
	"gvd_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.MySQL.Host == "" {
		global.Log.Fatal("未配置mysql")
		return nil
	}
	dsn := global.Config.MySQL.Dsn()
	var loglevel logger.LogLevel

	switch global.Config.MySQL.LogLevel {
	case "info":
		loglevel = logger.Info
	case "warn":
		loglevel = logger.Warn
	default:
		loglevel = logger.Error
	}

	mysqlLogger := logger.Default.LogMode(loglevel)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   mysqlLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		global.Log.Fatalln(fmt.Sprintf("[%s] mysql 连接失败,err:%s", dsn, err.Error()))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.MySQL.SetMaxIdleConns) // 最大连接数
	sqlDB.SetMaxOpenConns(global.Config.MySQL.SetMaxOpenConns) // 最多容量
	sqlDB.SetConnMaxLifetime(time.Hour * 4)                    //连接最大复用时间
	fmt.Printf("%s:%d mysql 初始化成功\n",global.Config.MySQL.Host,global.Config.MySQL.Port)
	return db
}
