package repositories

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mathegeapi/config"
	"time"
)

type Store struct {
	Config config.DatabaseConfig
	DB     *gorm.DB
}

func NewStore(config config.DatabaseConfig) *Store {
	return &Store{Config: config}
}

func (s *Store) Open() error {
	var err error

	s.DB, err = gorm.Open(postgres.Open(s.Config.DSN()), &gorm.Config{})

	if err != nil {
		return errors.New(fmt.Sprintf("Store.Open() err: %v", err))
	}

	db, err := s.DB.DB()

	if err != nil {
		return errors.New(fmt.Sprintf("Store.Open() err: %v", err))
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 5)

	return nil
}
