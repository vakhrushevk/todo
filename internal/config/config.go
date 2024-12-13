package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
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
			log.Printf("INFO help config %v", help)
			log.Fatalf("error with config, %v", err)
			// TODO:LOGGER
			// logger.fatal(err)
		}
	})
	return instance
}
