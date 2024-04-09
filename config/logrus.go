package config

type Logrus struct {
	AppName  string `yaml:"appname"`
	LogPath  string `yaml:"logpath"`
	NoData   bool   `yaml:"nodata"`
	NoErr    bool   `yaml:"noerr"`
	NoGlobal bool   `yaml:"noglobal"`
}
