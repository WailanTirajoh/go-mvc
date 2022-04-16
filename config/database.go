package config

import (
	"fmt"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	var DB *gorm.DB
	var err error

	config := NewConfig()

	switch connection := config.DB_CONNECTION; connection {
	case "mysql":
		DB, err = mysqlConnection(config)
		if err != nil {
			panic(err)
		}
	case "pgsql":
		DB, err = pgsqlConnection(config)
		if err != nil {
			panic(err)
		}
	default:
		DB, err = mysqlConnection(config)
		if err != nil {
			panic(err)
		}
	}

	return DB
}

func mysqlConnection(config VarConfig) (*gorm.DB, error) {
	DNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
	)

	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(
		model.User{},
	)

	return DB, nil
}

func pgsqlConnection(config VarConfig) (*gorm.DB, error) {
	panic("PGSQL connection is not ready yet.")
}
