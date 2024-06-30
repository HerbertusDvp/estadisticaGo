package main

import (
	"fmt"
)

func main() {

	datos := []float64{9, 7, 5, 8, 10}

	media := MedianaSlice(datos)
	//media := MedianaDatos(9, 7, 5, 8, 10)

	fmt.Println("Mediana: ", media)

}
