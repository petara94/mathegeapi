package services

import "C"
import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/stores"
)

type TaskService struct {
	CrudService[models.Task]
	rep *stores.TaskRepository
}

func NewTaskService(store *stores.Store) *TaskService {
	return &TaskService{CrudService: *NewCRUDService[models.Task](store), rep: stores.NewTaskStore(store)}
}

func (s *TaskService) Sub(r gin.IRouter) {
	s.CrudService.Sub(r)
}
