package config

type JWT struct {
	Expires int    `yaml:"expires"`
	Issuer  string `yaml:"issuer"`
	Secret  string `yaml:"secret"`
}
