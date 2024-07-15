package main

import (
	"fmt"
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

func TablaFrecuenciasAgrupados(datos [][]float64) [][]float64 {

	tabla := [][]float64{}
	clase := []float64{}

	minmax := MinMaxMatriz(datos)

	rango := minmax[1] - minmax[0]
	n := float64(NumeroDeDatos(datos)) // Numero de datos
	clases := 1 + 3.322 + math.Log10(n)
	k := Sturges(clases) // Numero de intervalos
	A := rango / float64(k)
	A = math.Ceil(A) //Amplitud de clase
	f := 0.0

	limiteI := minmax[0]
	limiteS := 0.0

	for i := 0; i < k; i++ {
		limiteS = limiteI + A

		clase = append(clase, limiteI)
		clase = append(clase, limiteS)
		clase = append(clase, (limiteI+limiteS)/2)

		for i := 0; i < len(datos); i++ {
			for j := 0; j < len(datos[i]); j++ {
				if datos[i][j] >= limiteI && datos[i][j] < limiteS {
					f += 1
				}
			}
		}

		clase = append(clase, f)
		if i == 0 {
			clase = append(clase, f)
		} else {
			clase = append(clase, f+tabla[i-1][4])
		}

		tabla = append(tabla, clase)
		clase = []float64{}
		limiteI = limiteS
		f = 0
	}

	return tabla
}

func MediaTabla(tabla [][]float64) float64 {

	fx := 0.0

	for i := 0; i < len(tabla); i++ {
		fx += tabla[i][2] * tabla[i][3]
	}
	n := tabla[len(tabla)-1][4]
	media := fx / float64(n)

	return media
}

func MedianaTabla(tabla [][]float64) float64 {

	n := tabla[len(tabla)-1][4]
	n2 := float64(n) / 2
	fila := 0
	mediana := 0.0

	for k, v := range tabla {
		fila = k
		if n2 < v[4] {
			break
		}
	}
	a := tabla[0][1] - tabla[0][0]
	mediana = tabla[fila][0] + ((n2-tabla[fila-1][4])/tabla[fila][3])*a

	mediana = Redondea(mediana, 2)
	return mediana
}

func ModaTabla(tabla [][]float64) {
	fila := 0
	aux := 0.0
	a := tabla[0][1] - tabla[0][0]

	for k, v := range tabla {
		if aux < v[3] {
			aux = v[3]
			fila = k
		}
	}
	resta1 := tabla[fila][3] - tabla[fila-1][3]
	resta2 := tabla[fila][3] - tabla[fila+1][3]
	moda := tabla[fila][0] + (resta1/(resta1+resta2))*a

	moda = Redondea(moda, 2)
	fmt.Println(moda)
}

func DispersionTabla(tabla [][]float64) []float64 {

	dispersion := []float64{}
	media := MediaTabla(tabla)
	sum := 0.0

	for _, v := range tabla {
		sum += math.Pow(v[2]-media, 2)
	}

	varianza := sum / tabla[len(tabla)-1][4]
	varianza = Redondea(varianza, 2)
	desviacionE := Redondea(math.Sqrt(varianza), 2)
	cv := Redondea(desviacionE/media, 2)

	dispersion = append(dispersion, varianza)
	dispersion = append(dispersion, desviacionE)
	dispersion = append(dispersion, cv)

	return dispersion
}
