package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	// redis config
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	// mysql config
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var GlobalConfig *Config

func init() {
	conf := new(Config)
	yamlConfFile, err := os.ReadFile("F:\\workspace_go\\gin-demo\\config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = yaml.Unmarshal(yamlConfFile, conf)
	if err != nil {
		log.Fatalln(err.Error())
	}
	GlobalConfig = conf
}
