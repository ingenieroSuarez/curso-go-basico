package pc

import "fmt"

type Pc struct {
	Ram       int
	Disk      int
	Brand     string
	monitores int // esto no se va a ver por que esta en minuscuula
}

func (myPc *Pc) DuplicarRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc *Pc) Ping() {
	fmt.Println(myPc)
}
