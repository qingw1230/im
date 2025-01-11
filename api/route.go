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

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/user/sendMsg", service.SendMsg)

	return r
}
