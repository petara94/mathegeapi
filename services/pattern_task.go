package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/stores"
)

type PatternTaskService struct {
	CrudService[models.PatternTask]
	rep stores.PatternTaskRepository
}

func NewPatternTaskService(store *stores.Store) *PatternTaskService {
	return &PatternTaskService{CrudService: *NewCRUDService[models.PatternTask](store), rep: *stores.NewPatternTaskStore(store)}
}

func (s *PatternTaskService) Sub(r gin.IRouter) {
	s.CrudService.Sub(r)
}
