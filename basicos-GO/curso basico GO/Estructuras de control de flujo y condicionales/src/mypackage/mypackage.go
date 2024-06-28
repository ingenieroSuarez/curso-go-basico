package mypackage

import "fmt"

type CarPublic struct { //las caracteristicas publicas llevan Mayus inicial
	Brand string
	Year  int
}

type CarPrivado struct { // las caracteristicas privadas son en minusculas
	brand string
	year  int
}

func imprimirSaludo(mensaje string) {
	fmt.Println("HOLA ", mensaje)
}

func ImprimirSaludoPublic(mensaje string) {
	fmt.Println("HOLA ", mensaje)
}
