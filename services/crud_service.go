package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/pkg/app"
	"mathegeapi/pkg/errors"
	"mathegeapi/pkg/utils"
	"mathegeapi/repositories"
	"net/http"
)

type CrudController[T any] struct {
	rep *repositories.CrudRepository[T, models.ID]
}

func NewCRUDController[T any](store *repositories.Store) *CrudController[T] {
	return &CrudController[T]{rep: repositories.NewCRUDRepository[T, models.ID](store)}
}

func (s *CrudController[T]) Sub(r gin.IRouter) {
	r.GET("/", s.GetAll)
	r.POST("/", s.Add)
	r.GET("/:id", s.Get)
	r.DELETE("/:id", s.Delete)
	r.PUT("/:id", s.UpdateUnsafe)
	r.PATCH("/:id", s.Update)
}

func (s *CrudController[T]) GetAll(c *gin.Context) {
	sender := app.Gin{C: c}

	entities := s.rep.GetAll()

	sender.Response(http.StatusOK, errors.SUCCESS, entities)
}

func (s *CrudController[T]) Get(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := utils.Atoi64(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	entity, err := s.rep.Get(models.ID(id))
	if err != nil {
		sender.Response(http.StatusNotFound, errors.ERROR_NOT_EXIST, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, entity)
}

func (s *CrudController[T]) Add(c *gin.Context) {
	sender := app.Gin{C: c}
	var json T

	err := sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	added, err := s.rep.Add(json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_ADD, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, added)
}

func (s *CrudController[T]) UpdateUnsafe(c *gin.Context) {
	sender := app.Gin{C: c}
	var json T

	id, err := utils.Atoi64(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	err = sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	updated, err := s.rep.UpdateUnsafe(models.ID(id), json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_UPDATE, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, updated)
}

func (s *CrudController[T]) Update(c *gin.Context) {
	sender := app.Gin{C: c}
	var json T

	id, err := utils.Atoi64(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	err = sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	updated, err := s.rep.Update(models.ID(id), json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_UPDATE, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, updated)
}

func (s *CrudController[T]) Delete(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := utils.Atoi64(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	s.rep.Delete(models.ID(id))

	sender.Response(http.StatusOK, errors.SUCCESS, nil)
}
