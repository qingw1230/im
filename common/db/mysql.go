package db

import (
	"github.com/qingw1230/im/common/db/models"
	"github.com/qingw1230/im/common/utils"
)

func LoadMySQLTable() {
	utils.MySQLDB.AutoMigrate(&models.UserBasic{})
}
