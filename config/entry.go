package config

type Config struct {
	System System `yaml:"SYSTEM"`
	Logrus Logrus `yaml:"LOG"`
	MySQL  Mysql  `yaml:"MYSQL"`
	Redis  Redis  `yaml:"REDIS"`
}
