package controllers

import (
	"log"
	"myapp/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Sport{})
	DB = db
}
