package main

import (
	"fmt"
)

func Punteros() {

	var nombre1 string = "Alejandro"
	nombre2 := "Matias"

	fmt.Println(nombre1, &nombre1)
	fmt.Println(nombre2, &nombre2)
}
