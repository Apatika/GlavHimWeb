package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server server   `yaml:"server"`
		DB     database `yaml:"database"`
	}

	server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}

	database struct {
		URI string `yaml:"uri"`
	}
)

var Cfg Config

func New() error {
	buff, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Println(err)
		return err
	}
	if err := yaml.Unmarshal(buff, &Cfg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
