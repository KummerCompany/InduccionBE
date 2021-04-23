//Ejercicio 5 - Condicionales
//Author: Melina Griffo

package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		a string
		b int
	)

	fmt.Println("Ingrese su apellido: ")
	fmt.Scanf("%s/n", &a)

	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d/n", &b)

	lopez := strings.ToLower("Lopez")

	if b >= 18 && b <= 64 {
		fmt.Println("Usted se encuentra en la población económicamente activa")

	} else {

		fmt.Println("Usted NO se encuentra en la población económicamente activa")
	}

	if a == lopez {
		fmt.Println("Usted tiene una fortuna pendiente con el Rey de Nigeria")

	}
}
