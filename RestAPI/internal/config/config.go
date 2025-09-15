package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

// env-default: "production"
type Config struct { // Capital Config is use becasue  to make it public not private
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"https_server"`
	// struct embedding
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to configration file") //use command-line flag like go run main.go -config=config.yaml
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	//exsures file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}
	return &cfg
}
