package config

import (
	"fmt"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type VarConfig struct {
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

func SetupConnection() *gorm.DB {
	var DB *gorm.DB
	var err error

	config := initConfig()

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

func initConfig() VarConfig {
	return VarConfig{
		DB_CONNECTION: GetEnv("DB_CONNECTION", "mysql"),
		DB_HOST:       GetEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:       GetEnv("DB_PORT", "3306"),
		DB_DATABASE:   GetEnv("DB_DATABASE", "2022_godb"),
		DB_USERNAME:   GetEnv("DB_USERNAME", "root"),
		DB_PASSWORD:   GetEnv("DB_PASSWORD", ""),
	}
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
