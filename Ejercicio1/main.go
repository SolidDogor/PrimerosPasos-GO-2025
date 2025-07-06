package main

import (
	"fmt"
)

func main() {
	var suma, mult, div int
	suma, mult, div = Operaciones(16, 8)
	fmt.Println(suma)
	fmt.Println(mult)
	fmt.Println(div)

	MostrarCurriculum()

	fmt.Println(Factorial(5))
}
