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
		DB     database
	}

	server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}

	database struct {
		URI string
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
	Cfg.DB.URI = os.Getenv("DB_PATH")
	return nil
}
