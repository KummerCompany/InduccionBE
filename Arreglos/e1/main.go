//Ejercicio 1 - Arreglos
//Author: Melina Griffo

package main

import "fmt"

func main() {

	var nombre string

	fmt.Println("Ingrese su nombre")
	fmt.Scanln(&nombre)

	prim_letra := string(nombre[0])
	letra_comp := "s"

	if prim_letra > letra_comp {

		fmt.Println("La letra mas grande es:", prim_letra)

	} else {
		fmt.Println("La letra mas grande es:", letra_comp)
	}
}
