package repositories

import "mathegeapi/models"

type TaskImageRepository struct {
	CrudRepository[models.TaskImage, models.ID]
}

func NewTaskImageRepository(store *Store) *TaskImageRepository {
	return &TaskImageRepository{CrudRepository: *NewCRUDRepository[models.TaskImage, models.ID](store)}
}
