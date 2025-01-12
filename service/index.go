package service

import (
	"html/template"

	"github.com/gin-gonic/gin"
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
