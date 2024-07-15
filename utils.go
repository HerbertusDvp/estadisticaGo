package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func LeerCSV(file string) [][]string {
	f, err := os.Open(file)

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

func MatrizStringToFloat64(datos [][]string) [][]float64 {

	matriz := [][]float64{}
	fila := []float64{}

	for i := 0; i < len(datos); i++ {
		for j := 0; j < len(datos[i]); j++ {
			aux, err := strconv.ParseFloat(datos[i][j], 64)
			if err != nil {
				fmt.Println("Error de convetsion con el dato: ", err)
			}
			fila = append(fila, aux)
			aux = 0.0
		}
		matriz = append(matriz, fila)
		fila = []float64{}
	}
	return matriz
}

func MinMaxMatriz(datos [][]float64) []float64 {

	res := []float64{0.0, 0.0}

	for i := 0; i < len(datos); i++ {
		for j := 0; j < len(datos[i]); j++ {

			aux := datos[i][j]

			if i == 0 && j == 0 {
				res[0] = aux
				res[1] = aux
			} else {
				if res[0] > aux {
					res[0] = aux
				}
				if res[1] < aux {
					res[1] = aux
				}
			}
		}
	}
	return res
}

func Sturges(numero float64) int {
	aux := int(numero - (numero - math.Floor(numero)))

	if aux%2 == 0 {
		return aux + 1
	}
	return aux
}

func NumeroDeDatos(datos [][]float64) int {
	n := 0

	for i := 0; i < len(datos); i++ {
		for j := 0; j < len(datos[i]); j++ {
			n += 1
		}
	}
	return n
}

func PrintTabla(tabla [][]float64) {

	for _, v := range tabla {
		fmt.Println(v)
	}
	fmt.Println()
}
