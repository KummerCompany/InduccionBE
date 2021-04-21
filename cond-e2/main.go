//Ejercicio 2 - Condicionales
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var a int
	var b int
	var c int
	var d int

	fmt.Println("Ingrese cuatro numeros diferentes: ")
	fmt.Scanf("%d/n", &a)
	fmt.Scanf("%d/n", &b)
	fmt.Scanf("%d/n", &c)
	fmt.Scanf("%d/n", &d)

	if (a < b) && (a < c) && (a < d) {
		fmt.Println("El numero mas chico es: ", a)

	}

	if (b < a) && (b < c) && (b < d) {
		fmt.Println("El numero mas chico es: ", b)

	}

	if (c < a) && (c < b) && (c < d) {
		fmt.Println("El numero mas chico es: ", c)

	}

	if (d < a) && (d < b) && (d < c) {
		fmt.Println("El numero mas chico es: ", d)

	}

}
