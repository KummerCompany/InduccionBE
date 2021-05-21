//Ejercicio 11 - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var numero, divisores int

	fmt.Print("Ingrese un numero: ")
	fmt.Scanf("%d", &numero)

	for div := 1; div <= numero; div++ {

		if numero%div == 0 {
			divisores++
		}
	}

	if divisores == 2 {
		fmt.Println("Cantidad de divisores:", divisores)
		fmt.Println("¿Es un numero primo?: Si")
	} else {
		fmt.Println("Cantidad de divisores:", divisores)
		fmt.Println("¿Es un numero primo?: No")
	}
}
