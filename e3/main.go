//Ejercicio 3
//Author: Griffo Melina

package main

import "fmt"

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

	//Ab potencia// pow
	//Raíz de A // sqrt

}
