package routers

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/config"
	v1 "mathegeapi/routers/api/v1"
	"mathegeapi/stores"
)

func InitApiRouter(apiConfig config.ApiConfig, store *stores.Store) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/images/", apiConfig.Server.ImagesDir)

	apiGroup := r.Group("api/v1")

	v1.InitTaskRouter(apiGroup.Group("tasks"), store)

	return r
}
