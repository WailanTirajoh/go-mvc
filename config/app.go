package config

import "strconv"

type (
	AppConfig struct {
		APP_KEY   string
		APP_NAME  string
		APP_DEBUG bool
		APP_URL   string
		APP_PORT  int64
	}
)

func NewAppConfig() *AppConfig {
	APP_DEBUG, _ := strconv.ParseBool(GetEnv("APP_DEBUG", "true"))
	APP_PORT, _ := strconv.ParseInt(GetEnv("APP_PORT", "9000"), 10, 32)
	return &AppConfig{
		APP_KEY:   GetEnv("APP_KEY", "mysecretkey"),
		APP_NAME:  GetEnv("APP_NAME", "Golang"),
		APP_DEBUG: APP_DEBUG,
		APP_URL:   GetEnv("APP_URL", "localhost"),
		APP_PORT:  APP_PORT,
	}
}
