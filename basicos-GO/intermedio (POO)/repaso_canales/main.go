package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("...")
	canal := make(chan int)
	go doSomething(canal)
	<-canal // la ejecución se bloquea hasta que el canal devuelva el mensaje
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
