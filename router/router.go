package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	Router = gin.Default()
	initTestRouter()
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	ErrNo   string      `json:"errNo"`
}
