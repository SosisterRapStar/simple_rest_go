package simple_rest_go

import (
	"log"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

type (
	Server struct {
		Host string `yaml:"host" env-default:"0.0.0.0"`
		Port string `yaml:"port" env-required:"true"`
	}
	Metrics struct {
		MetricsPort string `yaml:"metrics_port" env-default:":8080"`
	}

	Postgres struct {
		Url      string `yaml:"url" env-required:"true"`
		MaxConns int    `yaml:"max_conns" env-default:"10"`
		MinConns int    `yaml:"min_conns" env-default:"3"`
	}

	Storage struct {
		Postgres `yaml:"postgres"`
	}

	Config struct {
		Env     string `yaml:"env" env-default:"local"`
		Server  `yaml:"http_server"`
		Storage `yaml:"storage"`
		Metrics `yaml:"metrics"`
	}
)

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path is not set")
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

// logger also should be passed to functions like a paramater
// no global loggers now
func MustSetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}),
		)
	}
	return log
}
