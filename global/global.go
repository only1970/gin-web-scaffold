package global

import (
	"gin-web-scaffold/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *logrus.Logger
	MySQL  *gorm.DB
	Redis  *redis.Client
)
