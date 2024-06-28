package main

import "fmt"

type Empleado struct {
	id       int
	name     string
	vacation bool
}

func (e *Empleado) SetId(id int) {
	e.id = id
}

func (e *Empleado) SetName(name string) {
	e.name = name
}

func (e *Empleado) GetId() int {
	return e.id
}

func (e *Empleado) GetName() string {
	return e.name
}

func newEmployee(id int, name string, vacation bool) *Empleado {
	return &Empleado{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Empleado{}
	fmt.Printf("%+v\n", e)
	e.SetId(3)
	e.SetName("Marco")
	fmt.Printf("%+v\n", e)

	//creamos una funcion constructora

	e4 := newEmployee(5, "Pepe", true)
	fmt.Printf("%+v\n", *e4)

	empleado2 := new(Empleado) // otra forma de crear
	fmt.Println(*empleado2)    // al usar new, me retorna un apuntador a empleado2
}
