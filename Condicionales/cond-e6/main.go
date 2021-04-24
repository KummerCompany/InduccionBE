//Ejercicio seis - Condicionales
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var a float64

	fmt.Println("Ingrese su primer numero")
	fmt.Scanln(&a)

	var b float64

	fmt.Println("Ingrese su segundo numero")
	fmt.Scanln(&b)

	if b == 0 {
		fmt.Println("Existe una Indeterminacion. No es posible realizar el cálculo")
	} else {

		c := a / b

		fmt.Println("A/B es igual a: ", c)

	}

}
