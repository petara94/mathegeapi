package services

import "C"
import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/repositories"
)

type TaskController struct {
	CrudController[models.Task]
	rep *repositories.TaskRepository
}

func NewTaskController(store *repositories.Store) *TaskController {
	return &TaskController{CrudController: *NewCRUDController[models.Task](store), rep: repositories.NewTaskStore(store)}
}

func (s *TaskController) Sub(r gin.IRouter) {
	s.CrudController.Sub(r)
}
