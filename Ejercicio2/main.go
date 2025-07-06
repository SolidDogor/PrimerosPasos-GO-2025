//Crear un menú con 3 opciones. Cualquier otra opcion = error
//1) En la primera opción debemos crear un bucle que recorra un array
//(array de varias dimensiones)
//2) En la segunda opción debemos comprobar si un número es primo
//3) En la tercera opción debemos salir del menú

package main

import (
	"fmt"
)

func main() {
	opcion := 0

	for opcion != 3 {

		Menu()
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var arreglo [10]int
			for i := 0; i < 10; i++ {
				arreglo[i] = i
				fmt.Println(arreglo[i])
			}
		case 2:
			var num int
			fmt.Println("Ingrese el numero a evaluar: ")
			fmt.Scan(&num)
			if Primo(num) > 2 {
				fmt.Println("El numero es primo")
			} else {
				fmt.Println("El numero no es primo")
			}
		case 3:
		default:
			fmt.Println("Opcion no disponible!")
		}
	}

}

func Menu() {
	fmt.Println("1. Mostrar array")
	fmt.Println("2. Comprobar si un numero es primo")
	fmt.Println("3. Salir del programa")
	fmt.Println("Ingrese una opcion: ")
}

func Primo(num int) uint64 {
	var cantidad int
	for i := num; i > 0; i-- {
		if num%i == 0 {
			cantidad++
		}
	}
	return 1
}
