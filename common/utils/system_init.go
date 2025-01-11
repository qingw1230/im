package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

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

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	MySQLDB = db
}
