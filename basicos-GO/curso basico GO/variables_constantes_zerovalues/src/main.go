package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14141592
	const pi2 = 3.16141592
	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas
	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna valores a variables vacías, en otros lenguajes se asigna null
	var a int     // Go asigna 0
	var b float64 // Go asigna 0
	var c string  // Go asigna ''
	var d bool    // Go asigna false
	var e *int
	fmt.Println(a, b, c, d, e)
	if a == 0 {
		fmt.Println("A= ES CERO.")
	}
	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El área del cuadrado es:", areaCuadrado)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10
	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	numeroPar := 2
	resultadoEsPar := numeroPar % 2
	fmt.Println("Es par = ", resultadoEsPar)
	numeroPar++
	resultadoEsPar = numeroPar % 2
	fmt.Println("no es par, incremento = ", resultadoEsPar)

}
