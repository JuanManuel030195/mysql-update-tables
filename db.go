package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDB(auth Auth) (*sql.DB, error) {
	var err error
	var db *sql.DB

	cfg := mysql.Config{
		User:                 auth.Username,
		Passwd:               auth.Password,
		Net:                  auth.Protocol,
		Addr:                 auth.Endpoint + ":" + auth.Port,
		DBName:               auth.Database,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
