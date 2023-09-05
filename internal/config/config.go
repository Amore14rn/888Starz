package config

import (
	"flag"
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
	PORT           string        `yaml:"port" env:"PORT"`
	ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT"`
	MaxHeaderBytes int           `yaml:"max_header_bytes" env:"MAX_HEADER_BYTES"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var (
	configPath string
	instance   *Config
	once       sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		var configFileName = "configs/config.yaml" // Relative path from the current working directory
		flag.StringVar(
			&configPath,
			FlagConfigPathName,
			configFileName,
			"this is app config file",
		)
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			// Construct the absolute path using the current working directory
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error getting current working directory: %v", err)
			}
			configPath = filepath.Join(currentDir, configFileName)
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "888Starz - test task"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
