//Ejercicio 3 - Condicionales
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var a string
	var b string

	fmt.Println("Ingrese dos nombres diferentes: ")
	fmt.Scanf("%s/n", &a)
	fmt.Scanf("%s/n", &b)

	if a > b {
		fmt.Println("El nombre mas grande es: ", a)
	} else {
		fmt.Println("El nombre mas grande es:", b)
	}

}
