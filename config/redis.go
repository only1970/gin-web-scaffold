package config

type Redis struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"` //库名
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolsize"`
}
