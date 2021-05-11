//Ejercicio uno - Bucles
//Author: Melina Griffo

package main

import (
	"fmt"
)

func main() {

	var str string

	fmt.Println("Ingrese una palabra:")
	fmt.Scanln(&str)
	contador := 0

	for _, runa := range str {
		value := string(runa)
		if value == "a" || value == "e" || value == "i" || value == "o" || value == "u" {
			contador++

		}

	}

	fmt.Println("La cantidad de vocales que tiene su palabra es:", contador)
}
