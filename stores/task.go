package stores

import (
	"mathegeapi/models"
)

type TaskRepository struct {
	CrudRepository[models.Task, models.ID]
}

func NewTaskStore(store *Store) *TaskRepository {
	return &TaskRepository{CrudRepository: *NewCRUDRepository[models.Task, models.ID](store)}
}
