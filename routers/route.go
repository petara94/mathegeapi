package routers

import (
	"github.com/gin-gonic/gin"
	"mathegeapi/config"
	"mathegeapi/stores"
)

func GetRouter(apiConfig config.ApiConfig, store *stores.Store) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/images/", apiConfig.Server.ImagesDir)

	return r
}
