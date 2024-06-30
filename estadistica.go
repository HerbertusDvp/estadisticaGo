package main

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
