package main

import (
	"fmt"
)

func main() {

	var a float64
	fmt.Println("Ingrese su primer numero")
	fmt.Scanln(&a) // punteros y direcciones

	var b float64
	fmt.Println("Ingrese su segundo numero")
	fmt.Scanln(&b) // punteros y direcciones
	add := a + b

	fmt.Println("Su resultado es ", add)
}
