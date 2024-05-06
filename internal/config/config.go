package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server   server   `yaml:"server"`
		DB       database `yaml:"database"`
		Frontend frontend `yaml:"frontend"`
	}

	server struct {
		URI          string        `yaml:"uri"`
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
	}

	database struct {
		URI string `yaml:"uri"`
	}

	frontend struct {
		RefreshRate int `yaml:"refreshRate"`
	}
)

var Cfg Config
var DefaultCfg Config = Config{
	Server: server{
		URI:  "http://localhost",
		Port: "8080",
	},
	DB: database{
		URI: "mongodb://localhost:27017",
	},
	Frontend: frontend{
		RefreshRate: 10000,
	},
}

func SetDefaultEnv() {
	os.Setenv("CONFIG_PATH", "./config/config.yaml")
	os.Setenv("BUILD_PATH", "./bin")
	os.Setenv("APP_NAME", "glavhim")
	os.Setenv("STATIC_SOURCE_PATH", "./frontend/svelte/dist")
	os.Setenv("FRONTEND_PATH", "./frontend/svelte")
}

func New() {
	buff, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Println(err)
		log.Print("load default config")
		Cfg = DefaultCfg
	} else if err = yaml.Unmarshal(buff, &Cfg); err != nil {
		log.Println(err)
		log.Print("load default config")
		Cfg = DefaultCfg
	}
}
