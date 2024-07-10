package main

import (
	"math"
	"sort"
)

func Media(datos []float64, nDigitos float64) float64 {
	if len(datos) == 0 {
		return 0
	}
	suma := 0.0

	for _, dato := range datos {
		suma += dato
	}
	return Redondea(suma/float64(len(datos)), nDigitos)
}

func Mediana(datos []float64) float64 {
	sort.Float64s(datos)
	size := len(datos)
	indice := int(size / 2)

	if size == 0 {
		return 0
	}
	if size%2 == 0 {
		return (datos[indice] + datos[indice-1]) / 2
	} else {
		return datos[indice]
	}
}

func Moda(datos []float64) []float64 {

	valores := getDatosUnicos(datos)
	multimoda := MapMaxDato(valores)

	return multimoda
}

func getDatosUnicos(datos []float64) map[float64]int {
	valores := map[float64]int{}
	existe := false

	if !(len(datos) > 0) {
		return valores
	}

	for _, dato := range datos {

		if len(valores) > 0 {
			for key := range valores {
				if key == dato {
					existe = true
				}
			}

			if existe {
				valores[dato] = valores[dato] + 1
			} else {
				valores[dato] = 1
			}

		} else {
			valores[dato] = 1
		}
	}
	return valores
}

func MapMaxDato(datos map[float64]int) []float64 {
	maxFrec := 0
	modaSlice := []float64{}
	if len(datos) > 0 {
		for key, frecuencia := range datos {
			if maxFrec < frecuencia {
				maxFrec = datos[key]
			}
		}
		for key, frec := range datos {
			if frec == maxFrec {
				modaSlice = append(modaSlice, key)
			}
		}
	}
	return modaSlice
}

func Dispersion(datos []float64, nDigitos float64) []float64 {
	sort.Float64s(datos)
	resultado := []float64{0, 0, 0, 0}

	if len(datos) < 2 {
		return resultado
	}

	media := Media(datos, nDigitos)
	suma := 0.0
	for _, v := range datos {
		suma += math.Pow(v-media, 2)
	}
	resultado[0] = datos[len(datos)-1] - datos[0]        // Rango
	resultado[1] = Redondea(suma/float64(len(datos)), 2) // Varianza
	resultado[2] = Redondea(math.Sqrt(resultado[1]), 2)  //Desviacion estándar
	resultado[3] = Redondea(resultado[2]/media, 2)       //Coeficiente de varacion

	return resultado
}
