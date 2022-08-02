package main

import (
	"database/sql"
	"strconv"
)

type Localidad struct {
	idLocalidad int
	nombre      string
	idMunicipio sql.NullInt32
}

func GetLocalidad(db *sql.DB, table string, column string, value string, idMunicipio int) (Localidad, error) {
	var localidad Localidad
	query := "SELECT * FROM " + table + " WHERE " + column + "='" + value + "' AND idMunicipio=" + strconv.Itoa(idMunicipio)
	rows, err := db.Query(query)
	if err != nil {
		return localidad, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&localidad.idLocalidad, &localidad.nombre, &localidad.idMunicipio)
		if err != nil {
			return localidad, err
		}
	}
	return localidad, nil
}

func LocalidadExiste(db *sql.DB, table string, column string, value string, idMunicipio int) (bool, error) {
	query := "SELECT * FROM " + table + " WHERE " + column + "='" + value + "' AND idMunicipio=" + strconv.Itoa(idMunicipio)

	rows, err := db.Query(query)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func CreateLocalidad(db *sql.DB, table string, nombre string, idMunicipio int) error {
	query := "INSERT INTO " + table + " (nombre, idMunicipio) VALUES ('" + nombre + "', " + strconv.Itoa(idMunicipio) + ")"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
