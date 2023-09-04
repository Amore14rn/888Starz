package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Config struct {
	IsDevelopment bool     `yaml:"is-development" env:"IS-DEVELOPMENT" env-default:"false"`
	Server        Server   `yaml:"server"`
	Postgres      Postgres `yaml:"postgres"`
}

type Server struct {
	HOST           string        `yaml:"host" env:"HOST"`
	PORT           int           `yaml:"port" env:"PORT"`
	ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT"`
	MaxHeaderBytes int           `yaml:"max_header_bytes" env:"MAX_HEADER_BYTES"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

const (
	EnvConfigPathName = "CONFIG-PATH"
)

var (
	configPath string
	instance   *Config
	once       sync.Once
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			// Use the current working directory as the base path for the config file
			basePath, err := os.Getwd()
			if err != nil {
				log.Fatal("Failed to get current working directory")
			}
			configPath = filepath.Join(basePath, "configs", "env.dev.yaml")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "888Starz - test task"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance, nil
}
