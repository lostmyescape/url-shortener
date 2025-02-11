package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env     string `yaml:"env"`
	Storage struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		SslMode  string `yaml:"sslmode"`
	} `yaml:"storage"`
}

func LoadConfig() *Config {
	var cfg Config
	err := cleanenv.ReadConfig("/home/vanya/urlProject/config/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("Ошибка чтения config.yaml: %v", err)
	}
	return &cfg
}
