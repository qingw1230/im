package test

import (
	"database/sql"
	"testing"

	"github.com/qingw1230/im/db/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	dsn := "root:qin1002.@tcp(127.0.0.1:13306)/?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	defer sqlDB.Close()

	createDbSql := "CREATE DATABASE IF NOT EXISTS im"
	_, err = sqlDB.Exec(createDbSql)
	if err != nil {
		panic("failed to create database")
	}

	dsn = "root:qin1002.@tcp(127.0.0.1:13306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "test"
	db.Create(user)
}
