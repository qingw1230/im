package service

import (
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/common/db/models"
)

func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, nil)
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, nil)
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/foot.html",
	)
	if err != nil {
		panic(err)
	}

	userID, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userID)
	user.Identity = token
	ind.Execute(c.Writer, user)
}
