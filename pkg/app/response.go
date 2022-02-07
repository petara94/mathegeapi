package app

import (
	"github.com/gin-gonic/gin"

	"mathegeapi/pkg/errors"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorData struct {
	Error string `json:"error"`
}

func NewErrorData(err error) *ErrorData {
	return &ErrorData{Error: err.Error()}
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  errors.GetMsg(errCode),
		Data: data,
	})
	return
}
