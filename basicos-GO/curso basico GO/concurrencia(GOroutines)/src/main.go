package main

import (
	"fmt"
	"sync"
)

func say(mensaje string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // espero para que sea la ultima que se ejecute
	fmt.Println(mensaje)
}

func decir(mensaje string, c chan<- string) { // al adicionar <- a la derecha de chan, indico que es entrada
	c <- mensaje // indico que voy a ingresar un dato en el canal
}

/*
LOS waitGroup no se usan. son complejos de mantener, es mejor usar chanels
donde si se púeden enviar información las goroutines
*/
func main() { // main corre adentro de una goroutine
	var waitGroup sync.WaitGroup // acumula las goroutines y las va liberando de a poco

	fmt.Println("Hello")
	waitGroup.Add(1)            // voy a agregar una gouroutine al waitGroup
	go say("World", &waitGroup) // al anteponer Go se ejecuta de manera concurrente, se ejecuta en otro hilo
	// &waitGroup es un puntero
	waitGroup.Wait() // indico que espere por la gourutine.

	go func(mensaje string) { // función anonima
		fmt.Println(mensaje)
	}("-BYE-")
	chanels()
	//time.Sleep(time.Second * 1) // demorar un segundo
}

// Channels: La forma de organizar las goroutines
func chanels() {
	c := make(chan string, 1) // 1 es la cantidad limite
	fmt.Println("Hola")
	go decir("Adios", c)
	fmt.Println(<-c)
}
