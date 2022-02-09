package services

import "C"
import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/pkg/app"
	"mathegeapi/pkg/errors"
	"mathegeapi/stores"
	"net/http"
	"strconv"
)

type TaskService struct {
	store *stores.TaskStore
}

func NewTaskService(store *stores.Store) *TaskService {
	return &TaskService{store: stores.NewTaskStore(store)}
}

func (s *TaskService) Get(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	task, err := s.store.Get(uint(id))
	if err != nil {
		sender.Response(http.StatusNotFound, errors.ERROR_NOT_EXIST_TASK, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, task)
}

func (s *TaskService) GetAll(c *gin.Context) {
	sender := app.Gin{C: c}

	tasks, err := s.store.GetAll()
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, tasks)
}

func (s *TaskService) Add(c *gin.Context) {
	sender := app.Gin{C: c}
	var json models.Task

	err := sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	added, err := s.store.Add(json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_ADD, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, added)
}

func (s *TaskService) Update(c *gin.Context) {
	sender := app.Gin{C: c}
	var json models.Task

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	err = sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	updated, err := s.store.Update(uint(id), json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_UPDATE, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, updated)
}

func (s *TaskService) Delete(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	s.store.Delete(uint(id))

	sender.Response(http.StatusOK, errors.SUCCESS, nil)
}

func (s *TaskService) Sub(r *gin.RouterGroup) {
	r.GET("/", s.GetAll)
	r.POST("/", s.Add)
	r.GET("/:id", s.Get)
	r.DELETE("/:id", s.Delete)
	r.PUT("/:id", s.Update)
}
