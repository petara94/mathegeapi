package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/repositories"
)

type PatternTaskController struct {
	CrudController[models.PatternTask]
	rep repositories.PatternTaskRepository
}

func NewPatternTaskController(store *repositories.Store) *PatternTaskController {
	return &PatternTaskController{CrudController: *NewCRUDController[models.PatternTask](store), rep: *repositories.NewPatternTaskStore(store)}
}

func (s *PatternTaskController) Sub(r gin.IRouter) {
	s.CrudController.Sub(r)
}
