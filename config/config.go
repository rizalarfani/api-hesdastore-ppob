package config

import (
	"hesdastore/api-ppob/common/helper"

	"github.com/joho/godotenv"
)

var Config AppConfig

type AppConfig struct {
	Port      int       `json:"port"`
	AppName   string    `json:"appName"`
	Database  Database  `json:"database"`
	Digiflazz Digiflazz `json:"digiflazz"`
	RabbitMQ  RabbitMQ  `json:"rabbitMq"`
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

type Digiflazz struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	ApiKey   string `json:"apiKey"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	VHost    string `json:"vhost"`
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

	rabbitMQ := RabbitMQ{
		Host:     helper.GetEnv("RABBITMQ_HOST"),
		Port:     helper.GetEnvInt("RABBITMQ_PORT"),
		Username: helper.GetEnv("RABBITMQ_USERNAME"),
		Password: helper.GetEnv("RABBITMQ_PASSWORD"),
		VHost:    helper.GetEnv("RABBITMQ_VHOST"),
	}

	config := AppConfig{
		Port:     port,
		AppName:  helper.GetEnv("APP_NAME"),
		Database: db,
		RabbitMQ: rabbitMQ,
	}

	return &config
}

func (c *AppConfig) WithDigiflazz(d Digiflazz) {
	c.Digiflazz = d
}
