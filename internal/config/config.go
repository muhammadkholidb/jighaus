package config

import (
	"os"

	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
)

type ConfigData struct {
	Port string `yaml:"port"`
}

var Data *ConfigData

func Load() error {
	return LoadWithPath("./config.yml")
}

func LoadWithPath(path string) error {
	if Data != nil {
		log.Info("Config already loaded")
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&Data); err != nil {
		return err
	}
	return nil
}
