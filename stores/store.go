package stores

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mathegeapi/config"
	"mathegeapi/models"
	"sync"
)

type CRUDStore interface {
	Get(id uint, obj interface{}) error
	GetAll(objSlice interface{}) error
	Add(obj interface{}) error
	Update(id uint, obj models.AllowedToUpdate) error
	Delete(obj interface{}, id uint)
}

type Store struct {
	sync.RWMutex
	Config *config.DatabaseConfig
	DB     *gorm.DB
}

func (s *Store) Get(id uint, obj interface{}) error {
	s.RLock()
	defer s.RUnlock()

	res := s.DB.Preload(clause.Associations).Find(obj, id)
	if res.Error != nil {
		return errors.New("bad obj type")
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (s *Store) GetAll(objSlice interface{}) error {
	s.RLock()
	defer s.RUnlock()

	res := s.DB.Preload(clause.Associations).Find(objSlice)
	if res.Error != nil {
		return errors.New("bad obj type")
	}

	return nil
}

func (s *Store) Add(obj interface{}) error {
	s.Lock()
	defer s.Unlock()

	res := s.DB.Create(obj)
	if res.Error != nil {
		return errors.New(fmt.Sprintf("объект не может быть добавлен: %v", res.Error))
	}

	return nil
}

func (s *Store) Update(id uint, obj models.AllowedToUpdate) error {
	s.Lock()
	defer s.Unlock()

	res := s.DB.Model(obj).Where("id = ?", id).Omit("id", "created_at", "deleted_id").Updates(obj.Allowed())

	if res.Error != nil {
		return errors.New(fmt.Sprintf("объект не может быть обновлен: %v", res.Error))
	}

	s.DB.Preload(clause.Associations).Find(obj, id)

	return nil
}

func (s *Store) Delete(obj interface{}, id uint) {
	s.Lock()
	defer s.Unlock()

	s.DB.Delete(&obj, id)
}

func NewStore(config *config.DatabaseConfig) *Store {
	return &Store{Config: config}
}

func (s *Store) Open() (err error) {
	s.DB, err = gorm.Open(postgres.Open(s.Config.DSN()), &gorm.Config{})

	if err != nil {
		return errors.New(fmt.Sprintf("Store.Open() err: %v", err))
	}

	return
}
