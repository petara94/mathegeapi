package v1

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/services"
	"mathegeapi/stores"
)

func InitPatternTaskRouter(r *gin.RouterGroup, store *stores.Store) {
	taskService := services.NewPatternTaskService(store)

	r.GET("/", taskService.GetAll)
	r.POST("/", taskService.Add)
	r.GET("/:id", taskService.Get)
	r.DELETE("/:id", taskService.Delete)
	r.PUT("/:id", taskService.Update)
}
