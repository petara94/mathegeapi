package stores

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mathegeapi/config"
	"sync"
)

type Store struct {
	sync.RWMutex
	Config *config.DatabaseConfig
	DB     *gorm.DB
}

func NewStore(config *config.DatabaseConfig) *Store {
	return &Store{Config: config}
}

func (s *Store) Open() error {
	var err error
	s.DB, err = gorm.Open(postgres.Open(s.Config.DSN()), &gorm.Config{})

	if err != nil {
		return errors.New(fmt.Sprintf("Store.Open() err: %v", err))
	}

	return nil
}
