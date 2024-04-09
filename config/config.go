package config

import "fmt"

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	ENV  string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}
