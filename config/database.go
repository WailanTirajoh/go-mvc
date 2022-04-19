package config

import (
	"fmt"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

func NewDatabase() DbConfig {
	return DbConfig{
		DB_CONNECTION: GetEnv("DB_CONNECTION", "mysql"),
		DB_HOST:       GetEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:       GetEnv("DB_PORT", "3306"),
		DB_DATABASE:   GetEnv("DB_DATABASE", "2022_godb"),
		DB_USERNAME:   GetEnv("DB_USERNAME", "root"),
		DB_PASSWORD:   GetEnv("DB_PASSWORD", ""),
	}
}

func NewConnection() (*gorm.DB, error) {
	var DB *gorm.DB
	var err error

	dbConfig := NewDatabase()

	switch connection := dbConfig.DB_CONNECTION; connection {
	case "mysql":
		DB, err = mysqlConnection(dbConfig)
		if err != nil {
			return DB, err
		}
	case "pgsql":
		DB, err = pgsqlConnection(dbConfig)
		if err != nil {
			return DB, err
		}
	default:
		DB, err = mysqlConnection(dbConfig)
		if err != nil {
			return DB, err
		}
	}

	return DB, nil
}

func mysqlConnection(config DbConfig) (*gorm.DB, error) {
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

func pgsqlConnection(config DbConfig) (*gorm.DB, error) {
	panic("PGSQL connection is not ready yet.")
}
