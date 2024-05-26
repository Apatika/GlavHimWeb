package config

import (
	"errors"
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
		URI  string      `yaml:"uri"`
		Coll collections `yaml:"collections"`
	}

	collections struct {
		Orders    string `yaml:"orders"`
		Chemistry string `yaml:"chems"`
		Cargos    string `yaml:"cargos"`
		Customers string `yaml:"customers"`
		Users     string `yaml:"users"`
		City      string `yaml:"city"`
	}

	frontend struct {
		RefreshRate int `yaml:"refreshRate"`
	}
)

var Cfg Config
var DefaultCfg Config = Config{
	AppName: "glavhim",
	Server: server{
		URI:          "http://localhost",
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	},
	DB: database{
		URI: "mongodb://localhost:27017",
		Coll: collections{
			Orders:    "orders",
			Chemistry: "chems",
			Cargos:    "cargos",
			Customers: "customers",
			Users:     "users",
			City:      "cities",
		},
	},
	Frontend: frontend{
		RefreshRate: 10000,
	},
}

func New() {
	buff, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Println(err)
		if errors.Is(err, os.ErrNotExist) {
			log.Print("create config file")
			file, err := os.Create("./config/config.yaml")
			if err != nil {
				log.Panic("create config file error")
			}
			config, err := yaml.Marshal(DefaultCfg)
			if err != nil {
				log.Panic("create config file error")
			}
			buff = config
			_, err = file.Write(config)
			if err != nil {
				log.Panic("create config file error")
			}
		}
	}
	if err := yaml.Unmarshal(buff, &Cfg); err != nil {
		log.Println(err)
		log.Print("load default config")
		Cfg = DefaultCfg
	}
}
