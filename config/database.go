package config

import (
	"fmt"
	"log"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/spf13/viper"
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
		DB_CONNECTION: viperEnvVariable("DB_CONNECTION"),
		DB_HOST:       viperEnvVariable("DB_HOST"),
		DB_PORT:       viperEnvVariable("DB_PORT"),
		DB_DATABASE:   viperEnvVariable("DB_DATABASE"),
		DB_USERNAME:   viperEnvVariable("DB_USERNAME"),
		DB_PASSWORD:   viperEnvVariable("DB_PASSWORD"),
	}
}

func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
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
	panic("not yet ready")
}
