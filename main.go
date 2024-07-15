package main

import "fmt"

func main() {

	datos := LeerCSV("datos2.csv")
	matriz := MatrizStringToFloat64(datos)

	tabla := TablaFrecuenciasAgrupados(matriz)

	PrintTabla(tabla)

	fmt.Println(DispersionTabla(tabla))

}
