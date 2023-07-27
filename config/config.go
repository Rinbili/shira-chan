package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config = initConf()

type Yaml struct {
	Secret string `yaml:"secret"`
	Port   string `yaml:"port"`
	Rsakey string `yaml:"rsakey"`
	Ssl    struct {
		Enable bool   `yaml:"enable"`
		Crt    string `yaml:"crt"`
		Key    string `yaml:"key"`
	}
	Sql struct {
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Name   string `yaml:"name"`
	}
}

func initConf() *Yaml {
	config := new(Yaml)
	yamlFile, err := os.ReadFile("./config/config.yml")
	if err != nil {
		panic("config.yml not found")
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic("config.yml unmarshal fail")
	}
	return config
}
