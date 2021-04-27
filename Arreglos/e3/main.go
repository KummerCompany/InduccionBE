//Ejercicio tres - Arreglos
//Authot: Melina Griffo

package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string

	fmt.Println("Ingresa una palabra de máximo cinco caracteres")
	fmt.Scanln(&word)

	number_characters := int(len(word))
	max_characters := 5

	if number_characters != max_characters {
		fmt.Println("Dato mal ingresado")
		return
	} else {
		a_change := strings.Replace(word, "a", "X", 5)
		e_change := strings.Replace(a_change, "e", "X", 5)
		i_change := strings.Replace(e_change, "i", "X", 5)
		o_change := strings.Replace(i_change, "o", "X", 5)
		u_change := strings.Replace(o_change, "u", "X", 5)

		fmt.Println(u_change)

	}
}
