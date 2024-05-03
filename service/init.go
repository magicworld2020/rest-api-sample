package service

import (
	"fmt"
	"log"

	"github.com/magicworld2020/rest-api-sample/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbEngine *gorm.DB

func init() {
	dsn := "root:root@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DbEngine = db
	fmt.Println("init database ok")
}
