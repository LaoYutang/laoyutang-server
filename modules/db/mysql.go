package db

import (
	"fmt"
	"os"

	"github.com/laoyutang/laoyutang-server/modules/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Sql *gorm.DB

func initMySql() {
	// 获取连接串
	dsn := os.Getenv("LAOYUTANG_SQL") + "?parseTime=true"

	var err error
	Sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(">>> connect mysql failed, error=" + err.Error())
	}

	fmt.Println(">>> connect mysql success")
	// 设置数据库连接池参数
	sqlDB, _ := Sql.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数
	sqlDB.SetMaxIdleConns(10)

	// 建表
	migrateErr := Sql.AutoMigrate(&structs.User{}, &structs.Menu{})
	if migrateErr != nil {
		panic(">>> mysql migrate failed, error=" + migrateErr.Error())
	}
}
