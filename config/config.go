package config

import (
	"hesdastore/api-ppob/common/helper"

	"github.com/joho/godotenv"
)

var Config AppConfig

type AppConfig struct {
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `json:"maxOpenConnections"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnections    int    `json:"maxIdleConnections"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

func Load() *AppConfig {
	_ = godotenv.Load()

	port := helper.GetEnvInt("APP_PORT")

	db := Database{
		Host:                  helper.GetEnv("DB_HOST"),
		Port:                  helper.GetEnvInt("DB_PORT"),
		Name:                  helper.GetEnv("DB_NAME"),
		Username:              helper.GetEnv("DB_USER"),
		Password:              helper.GetEnv("DB_PASS"),
		MaxOpenConnections:    helper.GetEnvInt("DB_MAX_OPEN"),
		MaxLifeTimeConnection: helper.GetEnvInt("DB_CONN_MAX_LIFETIME_SEC"),
		MaxIdleConnections:    helper.GetEnvInt("DB_MAX_IDLE"),
		MaxIdleTime:           helper.GetEnvInt("DB_CONN_MAX_IDLE_SEC"),
	}

	config := AppConfig{
		Port:     port,
		Database: db,
	}
	return &config
}
