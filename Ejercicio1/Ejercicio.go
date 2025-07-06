//1) Definir una función a la que le lleguen como parámetros dos números y devuelva
//como resultado su multiplicación, su suma y su división
//2) Definir un struct para trabajar con una entidad que nos permita construir el
//curriculo de una persona. Nombre,email,dirección,tecnologías y experiencia en cada una
//de ellas y mostrar los datos por pantalla
//3) Función recursiva para el factorial de un número

package main

import (
	"fmt"
)

type Curriculum struct {
	edad        int
	nombre      string
	apellido    string
	email       string
	tecnologias [3]Tecnologias
}

type Tecnologias struct {
	nombre      string
	experiencia string
}

func Operaciones(num1 int, num2 int) (suma int, multiplicacion int, division int) {
	suma = num1 + num2
	multiplicacion = num1 * num2
	division = num1 / num2

	return suma, multiplicacion, division
}

func MostrarCurriculum() {

	var tecno [3]Tecnologias
	tecno[0].nombre = "c++"
	tecno[0].experiencia = "Intermedio"
	tecno[1].nombre = "java"
	tecno[1].experiencia = "Intermedio"
	tecno[2].nombre = "go"
	tecno[2].experiencia = "Junior"

	cur := Curriculum{21, "Alejandro", "Olivos", "olro@gmail.com", tecno}
	fmt.Println(cur)
}

func Factorial(num int) uint64 {
	if num >= 1 {
		return uint64(num) * Factorial(num-1)
	} else {
		return 1
	}

}
