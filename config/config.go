package config

import (
	"errors"
	"github.com/naoina/toml"
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

func LoadConfig(ConfigFile string) (cnf ApiConfig, err error) {
	f, err := os.Open(ConfigFile)

	if err != nil {
		return cnf, errors.New("Невозможно открыть файл конфигурации \"" + ConfigFilePath + "\"\n")
	}

	if err = toml.NewDecoder(f).Decode(&cnf); err != nil {
		return cnf, errors.New("Неизвестная ошибка")
	}

	return cnf, nil
}

func (dc DatabaseConfig) DSN() string {
	return "host=" + dc.Host + " " +
		"user=" + dc.Username + "  " +
		"password=" + dc.Password + " " +
		"dbname=" + dc.DBName + " " +
		"port=" + dc.Port + " " +
		"sslmode=disable "
}
