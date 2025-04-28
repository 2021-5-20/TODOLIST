package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"www.github.com/2021-5-20/t_d_list/models"
)

var JwtSecret = []byte("your_secret_key")
var db *gorm.DB

func InitDB() {
	dsn := "gormuser:gormpass@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	if err := db.AutoMigrate(&models.Memo{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	log.Println("数据库连接并迁移成功")
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("数据库未初始化")
	}
	return db
}
