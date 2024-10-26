package config

import (
	"github.com/04Akaps/gateway_module/log"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App []App `yaml:"app"`
}

type App struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Port    string `yaml:"port"`
	} `yaml:"app"`

	Producer Producer `yaml:"kafka"`
	Http     HttpCfg  `yaml:"http"`
}

func NewCfg(path string) Config {
	file, err := os.ReadFile(path)

	if err != nil {
		log.Log.Panic("Failed to open file", zap.String("path", path), zap.Error(err))
	}

	var c Config

	err = yaml.Unmarshal(file, &c)

	if err != nil {
		log.Log.Panic("Failed Unmarshal yaml file", zap.String("path", path), zap.Error(err))
	}

	return c
}
