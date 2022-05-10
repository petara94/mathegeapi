package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mathegeapi/config"
	"mathegeapi/models"
	"mathegeapi/services"
	"mathegeapi/stores"
	"net/http"
)

type IRouter interface {
	Group(path string) IRouter
	AddService(path string, service services.Service)
}

type Router struct {
	router *gin.RouterGroup
}

func (r Router) Group(path string) IRouter {
	return Router{router: r.router.Group(path)}
}

func (r Router) AddService(path string, service services.Service) {
	service.Sub(r.router.Group(path))
}

type App struct {
	config config.ApiConfig
	store  *stores.Store
	router *gin.Engine
}

func NewApp() *App {
	app := &App{}
	var err error

	app.config, err = config.LoadConfig(config.ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	app.store = stores.NewStore(app.config.Database)
	err = app.store.Open()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = app.store.DB.AutoMigrate(&models.Task{}, &models.TaskImage{}, &models.PatternTask{})
	if err != nil {
		log.Fatal(err.Error())
	}

	app.router = gin.New()
	app.router.Use(gin.Logger())
	app.router.Use(gin.Recovery())

	app.router.Static("/images", app.config.Server.ImagesDir)

	v1 := app.Group("/api/v1")
	v1.AddService("tasks", services.NewTaskService(app.store))
	v1.AddService("pattern_tasks", services.NewPatternTaskService(app.store))
	v1.AddService("task_images", services.NewTaskImageService(app.store))

	return app
}

func (a App) Group(path string) IRouter {
	return Router{router: a.router.Group(path)}
}

func (a App) AddService(path string, service services.Service) {
	service.Sub(a.router.Group(path))
}

func (a *App) Run() {
	err := http.ListenAndServe(":"+a.config.Server.Port, a.router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
