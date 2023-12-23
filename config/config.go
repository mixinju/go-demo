package config

import "gopkg.in/ini.v1"

var Config *ini.File

func Load() {
	var err error
	Config, err = ini.Load("")

	if err != nil {
		panic("加载配置失败")
	}
}

func Cos() *ini.Section {
	return Config.Section("cos")
}
