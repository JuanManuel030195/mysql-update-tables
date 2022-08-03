package main

import "database/sql"

type Municipio struct {
	idMunicipio int
	nombre      string
}

func GetMunicipio(db *sql.DB, table string, column string, value string) (Municipio, error) {
	var municipio Municipio
	query := "SELECT idMunicipio, nombre FROM " + table + " WHERE " + column + "='" + value + "'"
	rows, err := db.Query(query)
	if err != nil {
		return municipio, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&municipio.idMunicipio, &municipio.nombre)
		if err != nil {
			return municipio, err
		}
	}
	return municipio, nil
}

func GetMunicipios(db *sql.DB, table string) ([]Municipio, error) {
	var municipios []Municipio
	query := "SELECT idMunicipio, nombre FROM " + table
	rows, err := db.Query(query)
	if err != nil {
		return municipios, err
	}
	defer rows.Close()
	for rows.Next() {
		var municipio Municipio
		err = rows.Scan(&municipio.idMunicipio, &municipio.nombre)
		if err != nil {
			return municipios, err
		}
		municipios = append(municipios, municipio)
	}
	return municipios, nil
}

func MunicipioExiste(db *sql.DB, table string, column string, value string) (bool, error) {
	query := "SELECT * FROM " + table + " WHERE " + column + "='" + value + "'"

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

func CreateMunicipio(db *sql.DB, table string, municipio Municipio) error {
	query := "INSERT INTO " + table + " (nombre) VALUES ('" + municipio.nombre + "')"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
