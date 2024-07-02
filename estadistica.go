package main

import (
	"fmt"
	"sort"
)

func MediaDatos(datos ...float64) float64 {
	if len(datos) == 0 {
		return 0
	}
	suma := 0.0
	for _, dato := range datos {
		suma += dato
	}
	return suma / float64(len(datos))
}

func MediaSlice(datos []float64) float64 {
	if len(datos) == 0 {
		return 0
	}
	suma := 0.0

	for _, dato := range datos {
		suma += dato
	}
	return suma / float64(len(datos))
}

func MedianaDatos(datos ...float64) float64 {
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

func MedianaSlice(datos []float64) float64 {
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

func ModaDatos(datos ...float64) []float64 {

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
			for key, _ := range valores {
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
	fmt.Println(datos)
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
