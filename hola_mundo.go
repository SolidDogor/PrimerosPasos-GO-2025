package main

import "fmt"

func main() {

	//Comentario de una linea

	/*
	   Comentario
	   de
	   varias
	   lineas
	*/

	fmt.Println("Hola Mundo!")

	//Variables
	//Declaraación general: var nombre tipo

	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)

	//Variando una variable
	myString = "Y esta tambien es una cadena de texto!"
	fmt.Println(myString)

	//Se puede declarar sin especificar el tipo si se inicializa en la misma linea
	var myString2 = "Esta es otra cadena de texto"
	fmt.Println(myString2)

	//Trabajando con int (por defecto int = int32)
	var myInt int = 7
	fmt.Println(myInt)
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1) //El valor de myInt no cambia tras esto
	fmt.Println(myInt)

}
