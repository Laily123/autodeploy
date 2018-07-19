package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

var (
	Config = make(map[string]*ConfigStruct, 0)
)

type ConfigStruct struct {
	Name      string `toml:"name"`
	Secret    string `toml:"secret"`
	Dir       string `toml:"dir"`
	ShellName string `toml:"shell_name"`
}

type configsStruct struct {
	Configs []*ConfigStruct `toml:"project"`
}

func ParseConfig(filePath string) {
	var configs configsStruct
	_, err := toml.DecodeFile(filePath, &configs)
	if err != nil {
		log.Panic("parse config err: ", err)
	}
	for k, v := range configs.Configs {
		Config[v.Name] = configs.Configs[k]
	}
}
