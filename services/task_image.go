package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/stores"
)

type TaskImageService struct {
	CrudService[models.TaskImage]
	rep stores.TaskImageRepository
}

func NewTaskImageService(store *stores.Store) *TaskImageService {
	return &TaskImageService{CrudService: *NewCRUDService[models.TaskImage](store), rep: *stores.NewTaskImageRepository(store)}
}

func (s *TaskImageService) Sub(r gin.IRouter) {
	s.CrudService.Sub(r)
}
