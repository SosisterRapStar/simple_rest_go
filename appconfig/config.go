package appconfig

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Address string `yaml:"address" env-required:"true"`
}

type Postgres struct {
	Url      string `yaml:"url" env-required:"true"`
	MaxConns int    `yaml:"max_conns" env-default:"10"`
	MinConns int    `yaml:"min_conns" env-default:"3"`
}

type Storage struct {
	Postgres `yaml:"postgres"`
}

type Config struct {
	Env     string `yaml:"env" env-default:"local"`
	Server  `yaml:"http_server"`
	Storage `yaml:"storage"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s doesn't exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Can not read config")
	}

	return &cfg
}
