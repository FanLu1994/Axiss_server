package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configFile []byte
var config *Config

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
	} `yaml:"mysql"`
}

func GetGlobalConfig() *Config {
	return config
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic("open config file fail:" + err.Error())
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
}
