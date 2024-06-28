package main

import (
	"fmt"
)

func main() { // main corre adentro de una goroutine
	canal := make(chan string, 2)
	canal <- "Mensaje1"
	canal <- "Mensaje2"
	// canal <- "Mensaje3" // si trato de inresar otro msj, lanza error: "fatal error: all goroutines are asleep - deadlock!""
	gourotinesEnCanal := len(canal)
	capacidadMaximaDelCanal := cap(canal)
	fmt.Println(gourotinesEnCanal, capacidadMaximaDelCanal)

	// Range y Clouse
	close(canal) // cierro el canal, no recibe más datos
	//canal <- "Mensaje3" // lanza error "exit status 2" por que ya esta cerrado el canal

	for mensajeEnCanal := range canal {
		fmt.Println(mensajeEnCanal)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("MENSAJE A", email1)
	go message("MENSAJE B", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("EMAIL RECIBIDO DE EMAIL1", m1)
		case m2 := <-email2:
			fmt.Println("EMAIL RECIBIDO DE EMAIL2", m2)
		}

	}
}

func message(mensaje string, canal chan string) {
	canal <- mensaje
}
