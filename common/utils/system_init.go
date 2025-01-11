package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MySQLDB *gorm.DB
	RedisDB *redis.Client
)

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

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	_, err := RedisDB.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect redis")
	}
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = RedisDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RedisDB.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return msg.Payload, err
}
