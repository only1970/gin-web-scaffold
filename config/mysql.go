package config

import "strconv"

type Mysql struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Config          string `yaml:"config"` //高级配置，如charset
	DB              string `yaml:"db"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	LogLevel        string `yaml:"logLevel"` //日志等级
	SetMaxIdleConns int    `yaml:"SetMaxIdleConns"`
	SetMaxOpenConns int    `yaml:"SetMaxOpenConns"`
	// SetConnMaxLifetime int    `yaml:"SetConnMaxLifetime"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
