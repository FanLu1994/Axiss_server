package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configFile []byte

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
	} `yaml:"mysql"`
}

func GetGlobalConfig() *Config {
	var config *Config
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic("open config file fail:" + err.Error())
	}
}
