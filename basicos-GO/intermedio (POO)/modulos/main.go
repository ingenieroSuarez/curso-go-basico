package main

import (
	"fmt"

	hellomod "github.com/donvito/hellomod"
	hellomodv2 "github.com/donvito/hellomod/v2"
)

func main() {
	fmt.Println("holas")
	hellomod.SayHello()
	hellomodv2.SayHello(" - JF - ")
}
