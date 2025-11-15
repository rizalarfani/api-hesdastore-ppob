package config

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDatabase(dbConfig Database) (*sqlx.DB, error) {
	cfg := dbConfig
	dbs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sqlx.Open("mysql", dbs)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(Config.Database.MaxOpenConnections)
	db.SetMaxIdleConns(Config.Database.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(Config.Database.MaxLifeTimeConnection) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(Config.Database.MaxIdleTime) * time.Second)

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
