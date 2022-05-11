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
		return cnf, fmt.Errorf("config.Setup() - неизвестная ошибка: %v", err)
	}

	dbEnv := DatabaseConfig{
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if dbEnv != (DatabaseConfig{}) {
		cnf.Database = dbEnv
	}

	return cnf, nil
}

func (dc DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dc.Username, dc.Password, dc.Host, dc.Port, dc.DBName)
}
