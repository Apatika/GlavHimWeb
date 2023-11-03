package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

var Config config

func init() {
	Config = unmarshal()
}

type (
	config struct {
		Server server
	}

	server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}
)

func unmarshal() config {
	var cfg config
	buff, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(buff, &cfg); err != nil {
		log.Fatal("Error: config unmarshalling error")
	}
	return cfg
}
