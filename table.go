package main

import (
	"database/sql"
)

func DisableOnUpdate(db *sql.DB, table string, column string) {
	query := "ALTER TABLE " + table + " CHANGE " + column + " " + column + " DATETIME NOT NULL default CURRENT_TIMESTAMP"
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func EnableOnUpdate(db *sql.DB, table string, column string) {
	query := "ALTER TABLE " + table + " CHANGE " + column + " " + column + " DATETIME NOT NULL default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
