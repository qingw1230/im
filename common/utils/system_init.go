package utils

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB *gorm.DB

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config mysql content:", viper.Get("mysql"))
}

func InitMySQL() {
	sqlDB, err := sql.Open("mysql", viper.GetString("mysql.dsn1"))
	if err != nil {
		panic("failed to connect database")
	}
	defer sqlDB.Close()

	createDbSql := "CREATE DATABASE IF NOT EXISTS im"
	_, err = sqlDB.Exec(createDbSql)
	if err != nil {
		panic("failed to create database")
	}

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	MySQLDB = db
}
