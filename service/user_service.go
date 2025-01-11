package service

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/im/common/db/models"
	"github.com/qingw1230/im/common/utils"
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

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已注册",
		})
		return
	}

	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}

	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 用户模块
// @param name query string true "用户名"
// @param password query string true "密码"
// @success 200 {string} json{"code", "message"}
// @Router /user/findUserByNameAndPwd [get]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	pwd := c.Query("password")
	user := models.FindUserByName(name)
	flag := utils.ValidPassword(pwd, user.Salt, user.Password)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名或密码不正确",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
