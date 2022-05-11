package repositories

import "mathegeapi/models"

type PatternTaskRepository struct {
	CrudRepository[models.PatternTask, models.ID]
}

func NewPatternTaskStore(store *Store) *PatternTaskRepository {
	return &PatternTaskRepository{CrudRepository: *NewCRUDRepository[models.PatternTask, models.ID](store)}
}
