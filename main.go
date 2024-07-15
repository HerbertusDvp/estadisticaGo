package main

import "fmt"

func main() {

	datos := LeerCSV("datos.csv")
	matriz := MatrizStringToFloat64(datos)

	tabla := TablaFrecuenciasAgrupados(matriz)
	PrintTabla(tabla)

	fmt.Println("Media: ", MediaTabla(tabla))
	fmt.Println("Mediana: ", MedianaTabla(tabla))
	fmt.Println("Moda: ", ModaTabla(tabla))
	fmt.Println("Varianza: ", DispersionTabla(tabla)[0])
	fmt.Println("Desv Est.: ", DispersionTabla(tabla)[1])
	fmt.Println("Coe. Var.: ", DispersionTabla(tabla)[2])

}
