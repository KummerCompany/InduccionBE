//Ejercicio dos - Bucles
//Author: Melina Griffo

package main

import "fmt"

func main() {
	for {
		var value int

		fmt.Println("Ingrese números enteros del 1 al 9:")
		fmt.Scanln(&value)

		if value > 9 || value < 1 {
			fmt.Println("Dato mal ingresado")
			break

		}
	}
}
