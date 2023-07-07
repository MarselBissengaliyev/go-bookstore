package config

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
