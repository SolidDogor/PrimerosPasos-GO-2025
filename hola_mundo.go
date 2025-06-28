package main

import (
	"container/list"
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

	//Arreglos

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	//myArray[3] = 4 -> ERROR!
	fmt.Println(myArray[2])
	//fmt.Println(myArray[3]) -> ERROR!

	//Map

	myMap := make(map[string]int)
	myMap["Alejandro"] = 21
	myMap["Peter"] = 52
	myMap["Adrian"] = 28
	fmt.Println(myMap)
	fmt.Println(myMap["Alejandro"])

	myMap2 := map[string]int{"Alejandro": 21, "Peter": 52, "Adrian": 28}
	fmt.Println(myMap2)

	//List

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles

	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	//Funcion

	fmt.Println(myFunction())

	//Estructura

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Alejandro", 21}
	fmt.Println(myStruct)
}

func myFunction() string {
	return "Funcion que devuelve string"
}
