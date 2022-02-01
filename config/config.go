package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type ApiConfig struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Port     string
	Username string
	Host     string
	DBName   string
	Password string
}

var Config ApiConfig

func Setup(ConfigFile string) {
	f, err := os.Open(ConfigFile)

	if err != nil {
		log.Fatal("config.Setup() - невозможно открыть файл конфигурации \"" + ConfigFile + "\"\n")
	}

	if err = toml.NewDecoder(f).Decode(&Config); err != nil {
		log.Fatalf("config.Setup() - неизвестная ошибка: %v", err)
	}
}

func (dc DatabaseConfig) DSN() string {
	return "host=" + dc.Host +
		" user=" + dc.Username +
		" password=" + dc.Password +
		" dbname=" + dc.DBName +
		" port=" + dc.Port +
		"sslmode=disable"
}
