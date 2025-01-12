package main

import (
	"github.com/qingw1230/im/common/db"
	"github.com/qingw1230/im/common/utils"
	"github.com/qingw1230/im/internal/api"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	db.LoadMySQLTable()
	utils.InitRedis()

	r := api.Router()
	r.Run()
}
