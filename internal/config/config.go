package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		AppName  string   `yaml:"appname"`
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
		Coll collections `yaml:"collections"`
		URI  string      `yaml:"uri"`
	}

	collections struct {
		Orders    string `yaml:"orders"`
		Chemistry string `yaml:"chems"`
		Cargos    string `yaml:"cargos"`
		Clients   string `yaml:"clients"`
		Users     string `yaml:"users"`
	}

	frontend struct {
		RefreshRate int `yaml:"refreshRate"`
	}
)

var Cfg Config
var DefaultCfg Config = Config{
	AppName: "glavhim",
	Server: server{
		URI:  "http://localhost",
		Port: "8080",
	},
	DB: database{
		URI: "mongodb://localhost:27017",
		Coll: collections{
			Orders:    "orders",
			Chemistry: "chems",
			Cargos:    "cargos",
			Clients:   "clients",
			Users:     "users",
		},
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
