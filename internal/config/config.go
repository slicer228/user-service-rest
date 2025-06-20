package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	HTTPServerConfig
	DBConfig
}

type HTTPServerConfig struct {
	Host    string        `env:"HOST"`
	Port    string        `env:"PORT"`
	Timeout time.Duration `env:"TIMEOUT"`
}

type DBConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

func MustLoadConfig() *Config {
	var err error

	pathDb := os.Getenv("CONFIG_PATH_DB")
	pathHttp := os.Getenv("CONFIG_PATH_HTTP")

	if pathHttp == "" {
		panic("CONFIG_PATH_HTTP environment variable not set")
	}
	if pathDb == "" {
		panic("CONFIG_PATH_DB environment variable not set")
	}

	var config Config
	err = cleanenv.ReadConfig(pathDb, &config)

	if err != nil {
		panic(err)
	}

	err = cleanenv.ReadConfig(pathHttp, &config)

	if err != nil {
		panic(err)
	}

	return &config
}
