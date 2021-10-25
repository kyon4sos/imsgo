package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"im/model"
	"log"
	"os"
)

var dbInstance *gorm.DB

func init() {
	initDb()
}

func initDb() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			//SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      true,         // 禁用彩色打印
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger:newLogger,
	})
	if err != nil {
		log.Panicf("数据库连接异常 %v \n", err.Error())
	}
	dbInstance = db
	autoMigrate()
}
func GetDb() *gorm.DB {
	if dbInstance == nil {
		initDb()
	}
	return dbInstance
}

func autoMigrate()  {
	err := dbInstance.AutoMigrate(
		&model.ChatUser{},
		&model.TextMessage{},
		&model.Message{},
		&model.Conversation{})
	if err != nil {
		log.Panicf("数据库迁移失败%v \n",err.Error())
		return
	}
}