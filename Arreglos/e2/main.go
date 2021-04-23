//Ejercicio dos - Arreglos
//Authot: Melina Griffo

package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string

	fmt.Println("Ingresa una palabra")
	fmt.Scanln(&word)

	p := strings.Replace(word, string(word[2]), "P", 1)

	fmt.Println(p)

}
