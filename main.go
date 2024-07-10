package main

import "fmt"

func main() {
	file := "datos.csv"
	datos := LeerCSV(file)
	datosN := MatrizStringToFloat64(datos)
	fmt.Println(datosN)
	fmt.Println("Min y max: ", MinMaxMatriz(datosN))

}
