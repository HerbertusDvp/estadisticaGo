package main

import (
	"fmt"
)

func main() {

	datos := []float64{9, 7, 5, 8, 9, 8, 10, 5, 5}

	fmt.Println(Media(datos, 2), Mediana(datos), Moda(datos))

	fmt.Println("Dispersion: ", Dispersion(datos, 2))

}
