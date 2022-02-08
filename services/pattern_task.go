package services

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/models"
	"mathegeapi/pkg/app"
	"mathegeapi/pkg/errors"
	"mathegeapi/stores"
	"net/http"
	"strconv"
)

type PatternTaskService struct {
	store *stores.PatternTaskStore
}

func NewPatternTaskService(store *stores.Store) *PatternTaskService {
	return &PatternTaskService{store: stores.NewPatternTaskStore(store)}
}

func (p *PatternTaskService) Get(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	patternTask, err := p.store.Get(uint(id))
	if err != nil {
		sender.Response(http.StatusNotFound, errors.ERROR_NOT_EXIST_TASK, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, patternTask)
}

func (p *PatternTaskService) GetAll(c *gin.Context) {
	sender := app.Gin{C: c}

	patternTasks, err := p.store.GetAll()
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, patternTasks)
}

func (p *PatternTaskService) Add(c *gin.Context) {
	sender := app.Gin{C: c}
	var json models.PatternTask

	err := sender.C.BindJSON(&json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_WRONG_JSON, app.NewErrorData(err))
		return
	}

	added, err := p.store.Add(json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_ADD, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, added)
}

func (p *PatternTaskService) Update(c *gin.Context) {
	sender := app.Gin{C: c}
	var json models.PatternTask

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

	updated, err := p.store.Update(uint(id), json)
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.ERROR_UPDATE, app.NewErrorData(err))
		return
	}

	sender.Response(http.StatusOK, errors.SUCCESS, updated)
}

func (p *PatternTaskService) Delete(c *gin.Context) {
	sender := app.Gin{C: c}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sender.Response(http.StatusBadRequest, errors.INVALID_PARAMS, app.NewErrorData(err))
		return
	}

	p.store.Delete(uint(id))

	sender.Response(http.StatusOK, errors.SUCCESS, nil)
}
