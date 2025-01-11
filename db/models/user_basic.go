package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
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
