package repositories

import (
	"fmt"
	"gorm.io/gorm/clause"
	"mathegeapi/models"
	"reflect"
)

type Repository struct {
	Store *Store
}

func NewRepository(store *Store) *Repository {
	return &Repository{Store: store}
}

type CrudRepository[T any, K comparable] struct {
	Repository
}

func NewCRUDRepository[T any, K comparable](store *Store) *CrudRepository[T, K] {
	return &CrudRepository[T, K]{Repository: *NewRepository(store)}
}

func (rep *CrudRepository[T, K]) Get(id K) (*T, error) {
	var entity *T = new(T)
	rep.Store.DB.Preload(clause.Associations).Find(entity, id)

	idField := models.GetIdFieldNameOfEntity(entity)
	act := reflect.ValueOf(entity).Elem().FieldByName(idField)
	def := reflect.ValueOf(*new(K))

	// Не найдено если ключ равен своему нулевому значению
	if fmt.Sprint(act) == fmt.Sprint(def) {
		return nil, fmt.Errorf("запись не найдена")
	}

	return entity, nil
}

func (rep *CrudRepository[T, K]) GetAll() (entities []T) {

	rep.Store.DB.Preload(clause.Associations).Find(&entities)

	return entities
}

func (rep *CrudRepository[T, K]) Add(entity T) (*T, error) {

	res := rep.Store.DB.Create(&entity)
	if res.Error != nil {
		return nil, fmt.Errorf("запись не может быть добавлена: %v", res.Error)
	}

	return &entity, nil
}

func (rep *CrudRepository[T, K]) Delete(id K) {

	rep.Store.DB.Delete(&models.Task{}, id)
}

// UpdateUnsafe изменяет все поля, вне зависимости от их значения, за
// исключением полей встроенных типов
func (rep *CrudRepository[T, K]) UpdateUnsafe(id K, entity T) (*T, error) {

	reflect.ValueOf(&entity).Elem().FieldByName(models.GetIdFieldNameOfEntity(&entity)).Set(reflect.ValueOf(id))

	res := rep.Store.DB.Model(&entity).Updates(models.AllowedToUpdate(&entity))
	if res.Error != nil {
		return nil, fmt.Errorf("запись не может быть обновлена: %v", res.Error)
	}

	rep.Store.DB.Preload(clause.Associations).Find(&entity, id)

	return &entity, nil
}

// Update Изменяет только значения, которые не равны своему
// нулевому значению, а также не игнорирует встроенные типы
func (rep *CrudRepository[T, K]) Update(id K, entity T) (*T, error) {

	reflect.ValueOf(&entity).Elem().FieldByName(models.GetIdFieldNameOfEntity(&entity)).Set(reflect.ValueOf(id))

	res := rep.Store.DB.Updates(entity)
	if res.Error != nil {
		return nil, fmt.Errorf("запись не может быть обновлена: %v", res.Error)
	}

	rep.Store.DB.Preload(clause.Associations).Find(&entity, id)

	return &entity, nil
}
