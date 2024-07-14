package main

import "fmt"

func main() {

	datos := LeerCSV("datos2.csv")
	matriz := MatrizStringToFloat64(datos)

	//fmt.Println(matriz)

	tabla := TablaFrecuenciasAgrupados(matriz)

	fmt.Println(tabla)

}
