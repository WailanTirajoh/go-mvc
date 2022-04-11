package config

import (
	"fmt"

	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupConnection() *gorm.DB {
	const DNS = "root:@tcp(127.0.0.1:3306)/2022_godb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	DB.AutoMigrate(
		model.User{},
	)

	return DB
}
