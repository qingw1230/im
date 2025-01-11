package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/db/models"
)

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string true "用户名"
// @param password query string true "密码"
// @param repassword query string true "确认密码"
// @success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}
