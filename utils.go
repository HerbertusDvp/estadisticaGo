package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
)

func Max(datos []float64) float64 {
	sort.Float64s(datos)

	if len(datos) > 0 {
		return datos[len(datos)-1]
	}

	return 0
}

func Min(datos []float64) float64 {
	sort.Float64s(datos)

	if len(datos) > 0 {
		return datos[0]
	}

	return 0
}

func Redondea(numero float64, digitos float64) float64 {

	nDigitos := math.Pow(10, digitos)
	resultado := math.Round(numero*nDigitos) / nDigitos

	return resultado
}

func LeerCSV() [][]string {
	f, err := os.Open("datos.csv")

	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
	}
	defer f.Close()

	lector := csv.NewReader(f)
	lector.Comma = ','          //Especifica separador de columnas
	lector.Comment = '#'        //Especifica simbolo de comentario
	lector.FieldsPerRecord = -1 //Numero de comlumnas: -1 si se desconoce

	datos, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el leer los datos")
	}
	return datos
}
