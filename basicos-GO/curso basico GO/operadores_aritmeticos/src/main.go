package main

import "fmt"

func main() {
	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10
	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	numeroPar := 3
	resultadoEsPar := numeroPar % 2
	fmt.Println(resultadoEsPar)

}
