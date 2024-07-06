package main

import (
	"math"
	"sort"
)

func MaxOfSlice(datos []float64) float64 {
	sort.Float64s(datos)

	if len(datos) > 0 {
		return datos[len(datos)-1]
	}

	return 0
}

func MinOfSlice(datos []float64) float64 {
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
