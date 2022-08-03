package main

import (
	"database/sql"
	"strconv"
)

type Activo struct {
	idActivo    int
	idMunicipio sql.NullInt32
	idLocalidad sql.NullInt32
	municipio   sql.NullString
	localidad   sql.NullString
	idEdificio  sql.NullInt32
}

func GetActivos(db *sql.DB) ([]Activo, error) {
	var activos []Activo
	query := "SELECT idActivo, idMunicipio, idLocalidad, municipio, localidad, idPiso FROM activos"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var activo Activo
		err = rows.Scan(&activo.idActivo, &activo.idMunicipio, &activo.idLocalidad, &activo.municipio, &activo.localidad, &activo.idEdificio)
		if err != nil {
			return nil, err
		}
		activos = append(activos, activo)
	}
	return activos, nil
}

func GetActivosPorMunicipio(db *sql.DB, municipio string) ([]Activo, error) {
	var activos []Activo
	query := "SELECT idActivo, idMunicipio, idLocalidad, municipio, localidad FROM activos WHERE municipio='" + municipio + "'"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var activo Activo
		err = rows.Scan(&activo.idActivo, &activo.idMunicipio, &activo.idLocalidad, &activo.municipio, &activo.localidad)
		if err != nil {
			return nil, err
		}
		activos = append(activos, activo)
	}
	return activos, nil
}

func GetActivosPorMunicipioLocalidad(db *sql.DB, idMunicipio int, idLocalidad int) ([]Activo, error) {
	var activos []Activo
	query := "SELECT idEdificio FROM activos WHERE idMunicipio=" + strconv.Itoa(idMunicipio) + " AND idLocalidad=" + strconv.Itoa(idLocalidad)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var activo Activo
		err = rows.Scan(&activo.idEdificio)
		if err != nil {
			return nil, err
		}
		activos = append(activos, activo)
	}
	return activos, nil
}

func UpdateActivoById(db *sql.DB, idActivo int, column string, value int) error {
	query := "UPDATE activos SET " + column + "=" + strconv.Itoa(value) + " WHERE idActivo=" + strconv.Itoa(idActivo)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
