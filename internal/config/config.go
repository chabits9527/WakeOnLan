package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port int    `yaml:"port"`
	MAC  string `yaml:"MAC"`
}

func Load() (appConfig Config) {
	var config Config
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("无法读取 config.yaml 配置文件")
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("无法解析 config.yaml 配置文件")
	}
	return config
}
