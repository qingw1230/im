package service

import (
	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
