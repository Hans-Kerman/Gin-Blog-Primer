package config

import (
	"fmt"
	"log"
	"time"

	"github.com/Hans-Kerman/GinBlogPrimer/backend/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() error {
	dbConfig := AppConfig.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s "+
		"port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("初始化数据库连接出现问题: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("连接到数据库对象出现问题: %v", err)
	}
	global.Db = db
	return nil
}
