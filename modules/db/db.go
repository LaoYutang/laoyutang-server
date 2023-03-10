package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// 获取连接串
	dsn := os.Getenv("LAOYUTANG_SQL")

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(">>> connect database failed, error=" + err.Error())
	}

	fmt.Println(">>> connect database success")
	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数
	sqlDB.SetMaxIdleConns(10)
}

func GetDB() *gorm.DB {
	return db
}
