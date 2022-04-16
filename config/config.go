package config

import (
	"log"

	"github.com/spf13/viper"
)

type VarConfig struct {
	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string
}

func NewConfig() VarConfig {
	return VarConfig{
		DB_CONNECTION: GetEnv("DB_CONNECTION", "mysql"),
		DB_HOST:       GetEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:       GetEnv("DB_PORT", "3306"),
		DB_DATABASE:   GetEnv("DB_DATABASE", "2022_godb"),
		DB_USERNAME:   GetEnv("DB_USERNAME", "root"),
		DB_PASSWORD:   GetEnv("DB_PASSWORD", ""),
	}
}

func GetEnv(key string, fallback string) string {
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
		return fallback
	}

	return value
}
