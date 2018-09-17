package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var (
	Config = make(map[string]*ConfigStruct, 0)
	App    *appConfig
)

type ConfigStruct struct {
	Addr      string `toml:"url"`
	Secret    string `toml:"secret"`
	Dir       string `toml:"dir"`
	ShellName string `toml:"shell_name"`
}

type appConfig struct {
	Port string `toml:"port"`
}

type configsStruct struct {
	App     *appConfig      `toml:"app"`
	Configs []*ConfigStruct `toml:"project"`
}

func ParseConfig(filePath string) {
	var configs configsStruct
	_, err := toml.DecodeFile(filePath, &configs)
	if err != nil {
		log.Panic("parse config err: ", err)
	}
	for k, v := range configs.Configs {
		Config[v.Addr] = configs.Configs[k]
	}
	App = configs.App
}
