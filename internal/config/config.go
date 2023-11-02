package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server struct {
			Port         string        `yaml:"port"`
			ReadTimeout  time.Duration `yaml:"readTimeout"`
			WriteTimeout time.Duration `yaml:"writeTimeout"`
		}
	}
)

func New() Config {
	return unmarshal()
}

func unmarshal() Config {
	var cfg Config
	buff, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(buff, &cfg); err != nil {
		log.Fatal("Error: config unmarshalling error")
	}
	return cfg
}
