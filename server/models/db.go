package models

import (
	"fmt"
	"log"
	"server/models/dao"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=123456 dbname=musicapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	fmt.Println("数据库连接成功")
	dao.SetDefault(Db)
}
