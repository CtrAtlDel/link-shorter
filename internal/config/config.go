package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"  env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-default:"../"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"10s"`
	Idle_timeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config{
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path not set")
	}
	
	// Проверяем что файл существует если нет падаем с ошибкой
	if _,err:= os.Stat(configPath); os.IsNotExist(err) { 
		log.Fatalf("Config path doesnt exist %s",configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath,&cfg); err != nil {
		log.Fatalf("Connot parse file config %s", err)
	}

	return &cfg
}
