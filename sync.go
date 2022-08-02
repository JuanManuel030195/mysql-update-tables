package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	programStartTime := time.Now()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	// Load config from .env file
	var LOCAL_DB_ENDPOINT string = os.Getenv("LOCAL_DB_ENDPOINT")
	var LOCAL_DB_USERNAME string = os.Getenv("LOCAL_DB_USERNAME")
	var LOCAL_DB_PASSWORD string = os.Getenv("LOCAL_DB_PASSWORD")
	var LOCAL_DB_DB_NAME string = os.Getenv("LOCAL_DB_DB_NAME")
	var LOCAL_DB_NET_PROTOCOL string = os.Getenv("LOCAL_DB_NET_PROTOCOL")
	var LOCAL_DB_PORT string = os.Getenv("LOCAL_DB_PORT")

	var localAuth = Auth{
		Endpoint: LOCAL_DB_ENDPOINT,
		Username: LOCAL_DB_USERNAME,
		Password: LOCAL_DB_PASSWORD,
		Database: LOCAL_DB_DB_NAME,
		Protocol: LOCAL_DB_NET_PROTOCOL,
		Port:     LOCAL_DB_PORT,
	}

	localDB, err := ConnectToDB(localAuth)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to local DB")

	// Obtiene todos los activos
	activos, err := GetActivos(localDB)
	if err != nil {
		panic(err.Error())
	}

	// Obtiene todos los municipios
	var municipios []string
	for _, activo := range activos {
		if activo.municipio.Valid {
			municipioYaAgregado := false

			for _, municipio := range municipios {
				if municipio == activo.municipio.String {
					municipioYaAgregado = true
				}
			}

			if !municipioYaAgregado {
				municipios = append(municipios, activo.municipio.String)
				fmt.Println("Municipio obtenido: " + activo.municipio.String)
			}
		}
	}

	// Guarda los municipios en la base de datos
	for _, municipio := range municipios {
		municipioYaEnDB, err := MunicipioExiste(localDB, "municipios", "nombre", municipio)

		if err != nil {
			panic(err.Error())
		}

		if !municipioYaEnDB {
			CreateMunicipio(localDB, "municipios", Municipio{nombre: municipio})
			fmt.Println("Municipio agregado a DB: " + municipio)
		}
	}

	// Guarda las localidades en la base de datos
	localidadesGuardadas := 0
	for _, municipio := range municipios {
		activos, err := GetActivosPorMunicipio(localDB, municipio)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Localidades obtenidas para municipio: " + municipio)

		var localidades []string
		for _, activo := range activos {
			if activo.localidad.Valid {
				localidadYaAgregada := false

				for _, localidad := range localidades {
					if localidad == activo.localidad.String {
						localidadYaAgregada = true
					}
				}

				if !localidadYaAgregada {
					localidades = append(localidades, activo.localidad.String)
					fmt.Println("Localidad obtenida: " + activo.localidad.String)
				}
			}
		}

		municipioEnDB, err := GetMunicipio(localDB, "municipios", "nombre", municipio)

		if err != nil {
			panic(err.Error())
		}

		for _, localidad := range localidades {
			localidadYaEnDB, err := LocalidadExiste(localDB, "localidades", "nombre", localidad, municipioEnDB.idMunicipio)

			if err != nil {
				panic(err.Error())
			}

			if !localidadYaEnDB {
				CreateLocalidad(localDB, "localidades", localidad, municipioEnDB.idMunicipio)
				localidadesGuardadas++
				fmt.Println("Localidad agregada a DB: " + localidad)
			}
		}
	}
	fmt.Printf("Localidades guardadas: %d\r\n", localidadesGuardadas)

	// Actualiza activos con idMunicipio y idLocalidad
	activosActualizados := 0
	for _, activo := range activos {
		if activo.municipio.Valid {
			municipioEnDB, err := GetMunicipio(localDB, "municipios", "nombre", activo.municipio.String)

			if err != nil {
				panic(err.Error())
			}

			err = UpdateActivoById(localDB, activo.idActivo, "idMunicipio", municipioEnDB.idMunicipio)

			if err != nil {
				panic(err.Error())
			}

			if activo.localidad.Valid {
				localidadEnDB, err := GetLocalidad(localDB, "localidades", "nombre", activo.localidad.String, municipioEnDB.idMunicipio)

				if err != nil {
					panic(err.Error())
				}

				err = UpdateActivoById(localDB, activo.idActivo, "idLocalidad", localidadEnDB.idLocalidad)

				if err != nil {
					panic(err.Error())
				}
			}
		}

		activosActualizados++
	}

	fmt.Printf("Activos actualizados: %d\r\n", activosActualizados)

	programEndTime := time.Now()
	programDuration := programEndTime.Sub(programStartTime)
	fmt.Println("Program duration: ", programDuration)
}
