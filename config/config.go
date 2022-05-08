package config

import (
	"errors"
	"fmt"
	"github.com/naoina/toml"
	"os"
)

type ApiConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Port     string
	Username string
	Host     string
	DBName   string
	Password string
}

type ServerConfig struct {
	DBFromEnv bool
	Port      string
	ImagesDir string
}

func LoadConfig(ConfigFile string) (ApiConfig, error) {
	var cnf ApiConfig

	f, err := os.Open(ConfigFile)

	if err != nil {
		return cnf, errors.New("config.Setup() - невозможно открыть файл конфигурации \"" + ConfigFile + "\"\n")
	}

	if err = toml.NewDecoder(f).Decode(&cnf); err != nil {
		return cnf, errors.New(fmt.Sprintf("config.Setup() - неизвестная ошибка: %v", err))
	}

	if cnf.Server.DBFromEnv {
		cnf.Database = DatabaseConfig{
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Host:     os.Getenv("DB_HOST"),
			DBName:   os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
		}
	}

	return cnf, nil
}

func (dc DatabaseConfig) DSN() string {
	return "host=" + dc.Host +
		" user=" + dc.Username +
		" password=" + dc.Password +
		" dbname=" + dc.DBName +
		" port=" + dc.Port +
		" sslmode=disable"
}
