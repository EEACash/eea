package controller

import (
	"eea/config"
	"eea/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open(config.Configs.Mysql), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	globalDB = db

	log.Println("db connected success")

	err = globalDB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("auto migrate success")
}

func GetDB() *gorm.DB {
	return globalDB.Session(&gorm.Session{})
}
