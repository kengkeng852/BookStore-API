package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "kengkeng852:kongkeng852@tcp(127.0.0.1:3306)/kengdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}

