package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server server
	}

	server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}
)

func New() (Config, error) {
	var cfg Config
	buff, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Println(err)
		return Config{}, err
	}
	if err := yaml.Unmarshal(buff, &cfg); err != nil {
		log.Println(err)
		return Config{}, err
	}
	return cfg, nil
}
