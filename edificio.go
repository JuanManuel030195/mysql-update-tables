package main

import (
	"database/sql"
	"strconv"
)

type Edificio struct {
	idEdificio  int
	nombre      string
	numero      sql.NullString
	idLocalidad sql.NullInt32
}

func UpdateEdificioById(db *sql.DB, idEdificio int, column string, value int) error {
	query := "UPDATE edificios SET " + column + "='" + strconv.Itoa(value) + "' WHERE idEdificio=" + strconv.Itoa(idEdificio)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
