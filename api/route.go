package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/index", service.GetIndex)

	return r
}
