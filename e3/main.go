//Ejercicio 3
//Author: Griffo Melina

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Println("Ingrese su primer numero")
	fmt.Scanln(&a) //punteros y direcciones

	var b float64

	fmt.Println("Ingrese su segundo numero")
	fmt.Scanln(&b)

	fmt.Println("A + B es igual a:", a+b)

	fmt.Println("A - B es igual a:", a-b)

	fmt.Println("B - A es igual a:", b-a)

	fmt.Println("A/B es igual a:", a/b)

	c := math.Pow(a, b)

	fmt.Println("A elevado a B es igual a:", c)

	d := math.Sqrt(a)

	fmt.Println("La raiz cuadrada de A es igua a:", d)

}
