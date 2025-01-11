package main

import (
	"github.com/qingw1230/im/api"
	"github.com/qingw1230/im/common/db"
	"github.com/qingw1230/im/common/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	db.LoadMySQLTable()

	r := api.Router()
	r.Run()
}
