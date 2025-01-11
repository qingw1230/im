package models

import (
	"github.com/qingw1230/im/common/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Salt          string
	Phone         string
	Email         string
	ClientIP      string
	ClientPort    string
	LoginTime     uint64
	LogoutTime    uint64
	HeartbeatTime uint64
	IsLogout      bool
	DeviceInfo    string
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.MySQLDB.Create(&user)
}

func FindUserByName(name string) *UserBasic {
	user := UserBasic{}
	utils.MySQLDB.Where("name = ?", name).First(&user)
	return &user
}
