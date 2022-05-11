package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/repositories"
)

type TaskImageController struct {
	CrudController[models.TaskImage]
	rep repositories.TaskImageRepository
}

func NewTaskImageController(store *repositories.Store) *TaskImageController {
	return &TaskImageController{CrudController: *NewCRUDController[models.TaskImage](store), rep: *repositories.NewTaskImageRepository(store)}
}

func (s *TaskImageController) Sub(r gin.IRouter) {
	s.CrudController.Sub(r)
}
