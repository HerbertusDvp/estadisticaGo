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
