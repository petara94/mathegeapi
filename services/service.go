package services

import "github.com/gin-gonic/gin"

type CRUD interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
