package services

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	Sub(r gin.IRouter)
}

type CRUD interface {
	Service
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Add(c *gin.Context)
	UpdateUnsafe(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
