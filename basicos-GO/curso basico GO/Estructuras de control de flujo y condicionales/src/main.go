package main

import (
	"fmt"
	carro "fundamentos/src/mypackage" // forma de crear un alias
	pc "fundamentos/src/pc"
	"strconv"
	"strings"
)

func main() {
	defer fmt.Println("ESTO SE IMPRIME AL FINALIZAR GO USANDO 'defer'")

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10
	areaCirculo := PI * radioCirculo * radioCirculo
	fmt.Println("El Area del Circulo es :", areaCirculo)

	numeroPar := 2
	var algo string = "hola mundo"
	resultadoEsPar := numeroPar % 2
	fmt.Println(resultadoEsPar)

	fmt.Printf("%v, se inserta la vble al principio ", algo)
	fmt.Printf("\nel tipo de variable es: %T. \n", numeroPar)

	resultado := auxiliar("HOLA SO", 4, 4)
	fmt.Println("RESULTADO: ", resultado)
	_, segundoReturn := retornaVarios(6) // retorna 2 valores, pero tomo uno usando  '_'
	fmt.Println(segundoReturn)           //30

	pruebas := false
	if pruebas {
		controlDeFlujo()
		ciclos()
		parseo()
		condicionales()
		switchFuntion(15)
		arreglosYslices()
		recorrerSlices()
		esPalindromo("AmoR a roma")
		llaveValorConMaps()
		estructs()
		punteros()
		stringers()
	}
	intercaces()
}

func auxiliar(mensaje string, numeroA, numeroB int) int { //buena practica si son del mismo tipo.
	fmt.Println(mensaje, " ", numeroA, numeroB)
	return numeroA * numeroB
}
func retornaVarios(a int) (c, d int) {
	c = 11
	return a + c, a * 5
}

func ciclos() {
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
	counter := 0
	for counter < 10 {
		fmt.Println("Counter => ", counter)
		counter++
	}

}

func condicionales() bool {
	valor1 := 1
	valor2 := 2
	if valor1 > valor2 || valor1 == valor2 {
		fmt.Println("Es mayor o igual.. ")
		return true
	} else {
		fmt.Println("no es mayor")
		return false
	}
}

func parseo() {
	value, errorParseo := strconv.Atoi("55A")
	if errorParseo != nil {
		fmt.Println(errorParseo)
	} else {
		fmt.Println("PARSEO: ", value)
	}
}

func switchFuntion(value int) {
	// modulo := value % 2   => lo puedo definir acá o directo en el switch:
	switch modulo := value % 2; modulo { // se usa cuando es una sola vble
	case 0:
		fmt.Println(value, " ES PAR")
	default:
		fmt.Println(value, " ES IMPAR")
	}
}

//keywords defer, break y continue

func controlDeFlujo() {
	defer fmt.Println("ESTO SE IMPRIME AL FINALIZAR GO. solo funciona en el main...")
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("controlDeFlujo, continua = ", i)
			continue
		}
		if i == 7 {
			fmt.Println("deteniendo en ", i)
			break
		}
		fmt.Println("controlDeFlujo ", i)
	}
}

func arreglosYslices() {
	// ARREGLOS: INMUTABLES, SLICE, MUTABLES
	arreglo := [4]int{0}
	arreglo[1] = 1
	fmt.Println(arreglo, len(arreglo), cap(arreglo)) // len es 4, lo demas Go lo pone en 0

	numerosSlice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(numerosSlice, len(numerosSlice), cap(numerosSlice))
	// metodos slice
	fmt.Println(numerosSlice[0])   // 10
	fmt.Println(numerosSlice[:3])  // [10 11 12] el ultimo elemento es excluyente (:3), no incluye el 13
	fmt.Println(numerosSlice[2:4]) // [12 13]
	fmt.Println(numerosSlice[4:])  // [14 15] primer elemento es inclusivo (4:)
	//append
	numerosSlice = append(numerosSlice, 16)
	fmt.Println(numerosSlice)

	nuevoSlice := []int{17, 18, 19}
	numerosSlice = append(numerosSlice, nuevoSlice...)
	fmt.Println(numerosSlice)
}

func recorrerSlices() {
	fmt.Println("recorrerSlices")
	slicePalabras := []string{"hola", "-que-", "hace\n"}
	for i := range slicePalabras {
		fmt.Print(slicePalabras[i])
	}
}

func esPalindromo(texto string) {
	texto = strings.ToLower(texto)
	var textoInvertido string
	for i := len(texto) - 1; i >= 0; i-- {
		textoInvertido += string(texto[i])
		fmt.Printf("\n => TIPO:  %T ", texto[i])
		fmt.Print(" => valor ASCCI [i] : ", texto[i])
	}
	if textoInvertido == texto {
		fmt.Println("\n '", texto, "': es palindromo")
	}
	fmt.Println(textoInvertido)
}

func llaveValorConMaps() { // OJO, NO SE ALMACENAN EN ORDEN !! son concurrentes
	edadNombres := make(map[string]int) // map[string]int => INDICA
	edadNombres["JUAN"] = 36
	edadNombres["ANDRES"] = 24
	fmt.Printf("%T edadNombres", edadNombres)
	fmt.Println("\n ")
	for indice, edad := range edadNombres {
		println(indice, edad)
	}

	edadJuan := edadNombres["JUAN"]
	fmt.Println("EDAD DE JUAN=> ", edadJuan)
	edadFelipe, EdadEncontrada := edadNombres["FELIPE"] // EdadEncontrada=false => no existe Felipe
	fmt.Println("edad encontrada? => ", EdadEncontrada, "; edad: ", edadFelipe)
}

func estructs() { // en go no usamos clases, usamos structs
	var carroDeJuan carro.CarPublic
	carroDeJuan.Brand = "mercedez"
	fmt.Println(carroDeJuan)
	carro.ImprimirSaludoPublic(" pero con mayuscula")
}

func punteros() {
	a := 50        // tipo int  valor=> 50
	punteroA := &a // tipo *int valor=> 0xc000086000 => es la dirección en memoria
	println(punteroA)
	myPC := pc.Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	myOtraPC := pc.Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	myPC.Ping()
	myPC.DuplicarRam()
	myPC.DuplicarRam()
	myPC.Ping()
	myOtraPC.Ping()
}

func stringers() bool { // simplemente es una concatenación en la impresion
	return true
}

func intercaces() {
	myInterface := []interface{}{"hola", 12, true, 4.90}
	fmt.Println(myInterface)
}

func concurrencia() { // repasar..

}
