package config

import (
	"log/slog"
	"sync"
	"todo/pkg/logger"
	"todo/pkg/logger/sl"

	"github.com/ilyakaznacheev/cleanenv"
)

// host=localhost port=5678 dbname=todo user=postgres password=qwerty sslmode=disable"
type Config struct {
	Pg struct {
		Dsn string `yaml:"dsn"`
	}
	Http Http
}
type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (h *Http) Address() string {
	// TODO: Продумать получше
	return h.Host + ":" + h.Port
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info("help config", slog.String("help", help))
			logger.Fatal("Ошибка инициализации конфига", sl.Err(err))
		}
	})
	return instance
}
