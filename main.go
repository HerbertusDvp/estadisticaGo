package main

import (
	"fmt"
)

func main() {

	//datos := []float64{9, 7, 5, 8, 10}

	//media := MedianaSlice(datos)
	moda := ModaDatos(9, 7, 5, 8, 9, 8, 10, 5, 5)

	fmt.Println("Mediana: ", moda)

}
