package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")

	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.POST("/user/register", service.UserRegister)
	r.GET("/toRegister", service.ToRegister)
	r.POST("/user/login", service.UserLogin)
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)

	return r
}
