package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/docs"
	"github.com/qingw1230/im/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/user/createUser", service.CreateUser)
	r.POST("/user/login", service.FindUserByNameAndPwd)
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)

	return r
}
