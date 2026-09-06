package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_commerce?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
