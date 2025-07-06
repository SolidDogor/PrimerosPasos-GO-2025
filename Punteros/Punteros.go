package main

import (
	"fmt"
)

func Punteros() {

	var nombre1 string = "Alejandro"
	nombre2 := "Matias"

	fmt.Println(nombre1, &nombre1)
	fmt.Println(nombre2, &nombre2)
	//Creamos una nueva variable que tendrá el mismo valor que nombre1
	//ocupamos 2 espacios de memoria con el mismo valor
	nombre3 := nombre1
	fmt.Println(nombre3, &nombre3)
	//Si queremos utilizar el mismo valor que nombre1 podemos
	//podemos utilizar un puntero que apunte a la dirección de memoria que nombre1
	nombre4 := &nombre1
	fmt.Println(nombre4, *nombre4)
}
