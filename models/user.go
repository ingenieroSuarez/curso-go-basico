package models

type User struct { // struct es el equivalente a clase, es decir que esta es la clase usuario
	Id       string `json:"id"`    // inicia en mayusculapara usarse afuera del paquete "models", pero se usa notación json para serializarlo a minus...
	Email    string `json:"email"` // ...se inicia en minuscula cuando se va a usar solo en el paquete "models"
	Password string `json:"password"`
}
