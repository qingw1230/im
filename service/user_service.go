package service

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/qingw1230/im/common/db/models"
	"github.com/qingw1230/im/common/utils"
)

type paramsUserRegister struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	Phone      string `json:"phone"`
}

func UserRegister(c *gin.Context) {
	params := paramsUserRegister{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	data := models.FindUserByPhone(params.Phone)
	if data.Phone != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已注册",
		})
		return
	}

	if params.Password != params.RePassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}

	user := models.UserBasic{}
	user.Name = params.Name
	user.Phone = params.Phone
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	user.Password = utils.MakePassword(params.Password, salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}

type paramsUserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	params := paramsUserLogin{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	user := models.FindUserByPhone(params.Phone)
	flag := utils.ValidPassword(params.Password, user.Salt, user.Password)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名或密码不正确",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    user,
		})
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("MsgHandler 发送失败", err)
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
