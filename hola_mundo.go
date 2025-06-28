package main

import (
	"fmt"
	"reflect"
)

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

	myString2 = "Mostrando un entero con una cadena "
	fmt.Println(myString2, myInt)

	//Mostrando el tipo de dato con la biblioteca reflect
	var myType = reflect.TypeOf(myInt)
	fmt.Println(myType)

	//Trabajando con flotantes
	var myFloat float64 = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	//Sumando un int convertido a float con otro float
	fmt.Println(float64(myInt) + myFloat)

	//Trabajando con booleanos
	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declarada e inicializada de manera abreviada
	myString3 := "Esto es ootra cadena de texto"
	fmt.Println(myString3)

	//Trabajando con constantes
	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	//Control de flujo

	myString = "Hola"
	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

}
