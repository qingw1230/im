package models

import (
	"github.com/qingw1230/im/common/utils"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId  uint // 谁的关系信息
	TargetId uint // 对应的谁 /群 ID
	Type     int  // 对应的类型  1好友  2群  3xx
	Desc     string
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.MySQLDB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.MySQLDB.Where("id in ?", objIds).Find(&users)
	return users
}
